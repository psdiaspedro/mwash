package seguranca

import "golang.org/x/crypto/bcrypt"

func GerarHash(senha string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func ChecaSenha(senha, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}