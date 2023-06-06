package encryption

import (
	"io"
	"reflect"
	"unsafe"

	"filippo.io/age"
	"github.com/fxamacker/cbor"
	"github.com/pkg/errors"
	"golang.org/x/crypto/chacha20poly1305"
)

// AgeEncryptor implements the Encryptor interface using age
type AgeEncryptor struct {
	rs        []age.Recipient
	lastState []byte
	target    io.WriteCloser
}

func NewAgeEncryptor(recipients []string) (Encryptor, error) {
	var rs []age.Recipient
	for _, recipient := range recipients {
		r, err := age.ParseX25519Recipient(recipient)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse recipient")
		}
		rs = append(rs, r)
	}
	return &AgeEncryptor{rs: rs}, nil
}

func (e *AgeEncryptor) LoadState(lastState []byte) error {
	e.lastState = lastState
	return nil
}

func (e *AgeEncryptor) GetState() ([]byte, error) {
	writerValue := reflect.ValueOf(e.target).Elem()
	unwritten := writerValue.FieldByName("unwritten").Bytes()
	buf := writerValue.FieldByName("buf").Bytes()
	nonce := writerValue.FieldByName("nonce").Bytes()
	aField := writerValue.FieldByName("a")
	aField = reflect.NewAt(aField.Type(), unsafe.Pointer(aField.UnsafeAddr())).Elem()
	aeadKey := reflect.ValueOf(aField.Interface()).Elem().FieldByName("key").Bytes()
	exported := WriterExported{
		A:            aeadKey,
		UnwrittenLen: len(unwritten),
		Buf:          buf,
		Nonce:        nonce,
	}
	return cbor.Marshal(exported, cbor.CanonicalEncOptions())
}

func (e *AgeEncryptor) Encrypt(in io.Reader, last bool) (io.ReadCloser, error) {
	var err error
	reader, writer := io.Pipe()
	if e.lastState == nil {
		go func() {
			e.target, err = age.Encrypt(writer, e.rs...)
			if err != nil {
				writer.CloseWithError(err)
				return
			}
			_, err = io.Copy(e.target, in)
			if err != nil {
				writer.CloseWithError(err)
				return
			}
			if last {
				// Write final encryption block
				e.target.Close()
			}
			// Close the pipe writer to signal the end of the stream
			writer.Close()
		}()
	} else {
		e.target, err = age.Encrypt(&noopWriter{}, e.rs...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create encrypt stream")
		}
		// The target is the age/internal/stream/Writer, let's overwrite all the fields using reflection
		var unmarshalled WriterExported
		err = cbor.Unmarshal(e.lastState, &unmarshalled)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal last state")
		}
		v := reflect.ValueOf(e.target).Elem()
		bufField := v.FieldByName("buf")
		bufField = reflect.NewAt(bufField.Type(), unsafe.Pointer(bufField.UnsafeAddr())).Elem()
		buf := *(*[EncChunkSize]byte)(unmarshalled.Buf)
		bufField.Set(reflect.ValueOf(interface{}(buf)))
		unwrittenField := v.FieldByName("unwritten")
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(unwrittenField.UnsafeAddr()))
		hdr.Data = bufField.UnsafeAddr()
		hdr.Len = unmarshalled.UnwrittenLen
		hdr.Cap = EncChunkSize
		nonceField := v.FieldByName("nonce")
		nonceField = reflect.NewAt(nonceField.Type(), unsafe.Pointer(nonceField.UnsafeAddr())).Elem()
		nonce := *(*[NonceSize]byte)(unmarshalled.Nonce)
		nonceField.Set(reflect.ValueOf(nonce))
		aead, err := chacha20poly1305.New(unmarshalled.A)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create chacha20poly1305")
		}
		aField := v.FieldByName("a")
		aField = reflect.NewAt(aField.Type(), unsafe.Pointer(aField.UnsafeAddr())).Elem()
		aField.Set(reflect.ValueOf(aead))
		dstField := v.FieldByName("dst")
		dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
		dstField.Set(reflect.ValueOf(writer))
		go func() {
			_, err := io.Copy(e.target, in)
			if err != nil {
				writer.CloseWithError(err)
			}
			if last {
				e.target.Close()
			}
			writer.Close()
		}()
	}

	return reader, nil
}

type noopWriter struct {
}

func (nw *noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

const (
	KeySize      = chacha20poly1305.KeySize
	NonceSize    = chacha20poly1305.NonceSize
	EncChunkSize = 65552
)

type WriterExported struct {
	A            []byte
	UnwrittenLen int
	Buf          []byte
	Nonce        []byte
}
