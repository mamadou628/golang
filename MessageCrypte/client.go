package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// Fonction de gestion des erreurs
func gestionErreur(err error) bool {
	if err != nil {
		fmt.Println("Erreur :", err)
		return true
	}
	return false
}

// Fonction pour charger une clé privée à partir d'un fichier
func loadPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	keyFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("le fichier ne contient pas une clé privée valide")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'analyse de la clé privée : %v", err)
	}

	return privateKey, nil
}

func main() {
	const (
		IP   = "127.0.0.1"
		PORT = "3569"
	)

	var wg sync.WaitGroup

	// Connexion au serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	if err != nil {
		fmt.Printf("Impossible de se connecter au serveur %s:%s : %v\n", IP, PORT, err)
		return
	}
	defer conn.Close()
	fmt.Println("Connecté au serveur.")

	wg.Add(2)

	// Goroutine pour envoyer des messages au serveur
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Entrez un message à envoyer au serveur : ")
			text, err := reader.ReadString('\n')
			if gestionErreur(err) {
				break
			}
			text = strings.TrimSpace(text)

			_, writeErr := conn.Write([]byte(text + "\n"))
			if gestionErreur(writeErr) {
				break
			}
		}
	}()

	// Goroutine pour recevoir des messages du serveur
	go func() {
		defer wg.Done()
		for {
			// Lire un message chiffré
			encryptedMessage, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Erreur lors de la réception du message ou connexion au serveur perdue :", err)
				break
			}
			encryptedMessage = strings.TrimSpace(encryptedMessage)
			fmt.Println("\nMessage chiffré reçu :\n", encryptedMessage)

			// Demander une clé privée pour déchiffrer
			decryptmessage(encryptedMessage)
		}
	}()

	// Attendre la fin des deux goroutines
	wg.Wait()
	fmt.Println("Connexion terminée.")
}

// Déchiffrer un message chiffré
func decryptmessage(encryptedMessage string) {
	// Demander une clé privée pour déchiffrer
	fmt.Print("Entrez le chemin de votre clé privée pour déchiffrer le message ou appuyez sur Entrée pour ignorer : ")
	reader := bufio.NewReader(os.Stdin)
	privateKeyPath, _ := reader.ReadString('\n')
	privateKeyPath = strings.TrimSpace(privateKeyPath)

	if privateKeyPath != "" {
		// Charger la clé privée
		fmt.Println("Chargement de la clé privée à partir de :", privateKeyPath)
		privateKey, err := loadPrivateKeyFromFile(privateKeyPath)
		if gestionErreur(err) {
			fmt.Println("Erreur lors du chargement de la clé privée.")
			return
		}

		// Décoder le message (Base64)
		fmt.Println("Décodage du message chiffré...")
		decodedMessage, err := base64.StdEncoding.DecodeString(encryptedMessage)
		if gestionErreur(err) {
			fmt.Println("Erreur lors du décodage du message.")
			return
		}

		// Afficher le message décodé pour vérifier
		fmt.Println("Message décodé : ", string(decodedMessage))

		// Déchiffrement avec le padding PKCS1v15
		fmt.Println("Déchiffrement du message avec PKCS1v15...")
		decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedMessage)
		if gestionErreur(err) {
			fmt.Println("Erreur de déchiffrement :", err)
			return
		}

		// Afficher le message déchiffré
		fmt.Println("Message déchiffré : ", string(decryptedMessage))
	} else {
		// Ignorer le déchiffrement
		fmt.Println("Vous avez choisi de ne pas déchiffrer ce message.")
	}
}
