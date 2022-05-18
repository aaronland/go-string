// package salt provides methods for generating string values that can be used as (hashing) salts.
package salt

import (
	"errors"
	"github.com/aaronland/go-string/random"
)

var min_length int

func init() {
	min_length = 32
}

// Salt is a struct containing methods and values that can be used as (hashing) salts.
type Salt struct {
	salt string
}

// String returns the string value of 's'.
func (s *Salt) String() string {
	return s.salt
}

// String returns the byte value of 's'.
func (s *Salt) Bytes() []byte {
	return []byte(s.salt)
}

// SaltOptions is a struct containing configuration options for the `NewRandomSalt` method.
type SaltOptions struct {
	// The length of the final salting string.
	Length int
	// A boolean flag indicating that the salt should only contain ASCII characters.
	ASCII bool
}

// DefaultOptions returns an `SaltOptions` instance with no limits or restrictions save a minimum length of 32 bytes.
func DefaultSaltOptions() *SaltOptions {

	opts := SaltOptions{
		Length: min_length,
		ASCII:  false,
	}

	return &opts
}

// NewRandomSalt return a new `Salt` instance for a random string configured by 'opts'.
func NewRandomSalt(opts *SaltOptions) (*Salt, error) {

	str_opts := random.DefaultOptions()
	str_opts.Length = opts.Length
	str_opts.ASCII = opts.ASCII

	s, err := random.String(str_opts)

	if err != nil {
		return nil, err
	}

	return NewSaltFromString(s)
}

// NewSaltFromString returns a new `Salt` instance derived from 's'.
func NewSaltFromString(s string) (*Salt, error) {

	_, err := IsValidSalt(s)

	salt := Salt{s}

	if err != nil {
		return nil, err
	}

	return &salt, nil
}

// IsValidSalt returns a boolean value indicating whether 's' can be used a salt.
func IsValidSalt(s string) (bool, error) {

	if len(s) < min_length {
		return false, errors.New("salt is too short")
	}

	return true, nil
}
