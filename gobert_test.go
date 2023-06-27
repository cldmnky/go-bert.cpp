package gobert_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/go-skynet/go-bert.cpp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("gobert binding", func() {
	Context("Declaration", func() {
		It("fails with no model", func() {
			model, err := New("not-existing")
			Expect(err).To(HaveOccurred())
			Expect(model).To(BeNil())
		})
	})
	Context("Embedding", func() {
		It("get embeddings", func() {
			model, err := New(filepath.Join("fixtures", "model.bin"))
			Expect(err).ToNot(HaveOccurred())
			Expect(model).ToNot(BeNil())
			embeddings, err := model.Embeddings("foo")
			Expect(err).ToNot(HaveOccurred())
			fmt.Println(embeddings)
		})
		It("truncate large context", func() {
			model, err := New(filepath.Join("fixtures", "model.bin"))
			Expect(err).ToNot(HaveOccurred())
			Expect(model).ToNot(BeNil())
			f, err := os.ReadFile("fixtures/shakespeare.txt")
			Expect(err).ToNot(HaveOccurred())
			embeddings, err := model.Embeddings(string(f))
			Expect(err).ToNot(HaveOccurred())
			fmt.Println(embeddings)
		})
	})
	Context("Tokenize", func() {
		It("get tokens", func() {
			model, err := New(filepath.Join("fixtures", "model.bin"))
			Expect(err).ToNot(HaveOccurred())
			Expect(model).ToNot(BeNil())
			f, err := os.ReadFile("fixtures/shakespeare.txt")
			Expect(err).ToNot(HaveOccurred())
			size, err := model.Tokenize(string(f))
			Expect(err).ToNot(HaveOccurred())
			Expect(size).To(Equal(512))
		})
	})
})
