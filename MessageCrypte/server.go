package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"sync"
)

func gestionerreur(err error) {
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

const (
	Ip   = "127.0.0.1"
	Port = "3569"
)

func main() {
	// Générer une paire de clés RSA

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	//privatekey, err := cle.generateRSAKeys()
	gestionerreur(err)
	publicKey := &privateKey.PublicKey
	// Écrire la clé publique dans un fichier
	err = savePublicKeyToFile("public_key.pem", publicKey)
	gestionerreur(err)
	fmt.Println("Clé publique enregistrée dans le fichier 'public_key.pem'")

	// Exporter la clé publique pour distribution
	publicKeyPEM := exportPublicKeyToPEM(publicKey)
	fmt.Println("Clé publique (PEM) :\n", publicKeyPEM)

	fmt.Println("Lancement du serveur...")

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", Ip, Port))
	gestionerreur(err)
	defer ln.Close()

	var clients []net.Conn      // Liste des connexions clients
	var clientsMutex sync.Mutex // Mutex pour synchroniser l'accès à la liste des clients

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erreur d'acceptation :", err)
			continue
		}

		clientsMutex.Lock()
		clients = append(clients, conn) // Ajouter le client à la liste
		clientsMutex.Unlock()

		fmt.Println("Un client est connecté depuis", conn.RemoteAddr())

		// Goroutine pour gérer les communications avec le client
		go func(c net.Conn) {
			defer func() {
				// Supprimer le client lors de sa déconnexion
				clientsMutex.Lock()
				for i, client := range clients {
					if client == c {
						clients = append(clients[:i], clients[i+1:]...)
						break
					}
				}
				clientsMutex.Unlock()
				fmt.Println("Client déconnecté :", c.RemoteAddr())
				c.Close()
			}()

			buf := bufio.NewReader(c)
			for {
				// Lire un message du client
				receivedMessage, err := buf.ReadString('\n')
				if err != nil {
					fmt.Println("Erreur lors de la lecture :", err)
					break
				}

				// Chiffrer le message avec la clé publique
				message := []byte(receivedMessage)
				encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
				if err != nil {
					fmt.Println("Erreur lors du chiffrement :", err)
					continue
				}

				// Convertir en Base64 pour transmission
				encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedMessage)

				fmt.Println("\nMessage reçu (clair) :", receivedMessage)
				fmt.Println("Message chiffré (Base64) :", encryptedBase64)

				// Diffuser le message chiffré à tous les autres clients
				clientsMutex.Lock()
				for _, client := range clients {
					if client != c { // Ne pas renvoyer au même client
						_, writeErr := client.Write([]byte(encryptedBase64 + "\n"))
						if writeErr != nil {
							fmt.Println("Erreur lors de l'écriture :", writeErr)
						}
					}
				}
				clientsMutex.Unlock()
			}
		}(conn)
	}
}

// Fonction pour exporter une clé publique au format PEM
func exportPublicKeyToPEM(publicKey *rsa.PublicKey) string {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de la clé publique :", err)
		return ""
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(publicKeyPEM)
}

func savePublicKeyToFile(filename string, publicKey *rsa.PublicKey) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage de la clé publique : %v", err)
	}
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return os.WriteFile(filename, pemData, 0644)
}
