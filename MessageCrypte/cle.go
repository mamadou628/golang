package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Génère une paire de clés RSA et les sauvegarde dans des fichiers
func generateRSAKeys() error {
	// Générer une clé privée RSA de 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("erreur lors de la génération de la clé privée : %v", err)
	}

	// Sauvegarder la clé privée dans un fichier
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier de clé privée : %v", err)
	}
	defer privateKeyFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage de la clé privée : %v", err)
	}

	// Générer la clé publique à partir de la clé privée
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier de clé publique : %v", err)
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("erreur lors de la génération de la clé publique : %v", err)
	}

	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage de la clé publique : %v", err)
	}

	fmt.Println("Clé privée et clé publique générées avec succès.")
	return nil
}

func main() {
	// Générer les clés RSA
	err := generateRSAKeys()
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Les clés ont été générées.")
	}
}
