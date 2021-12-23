package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
)

type labelRequest struct {
	Title   string
	BaseImg string
	Font    Font
}

type Font struct {
	Size float64
}

func (lr Font) fontSize() float64 {
	if lr.Size != 0 {
		return lr.Size
	}

	return 40
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	req := labelRequest{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot read the body")
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid JSON provided")
		return
	}

	decodedImg, err := base64.StdEncoding.DecodeString(req.BaseImg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot decode base64: %s", err)
		return
	}

	baseImg, _, err := image.Decode(bytes.NewReader(decodedImg))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot read the image: %s", err)
		return
	}

	img, err := drawLabel(baseImg, req.Title, req.Font.fontSize())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot draw the label: %s", err)
		return
	}

	w.Header().Set("content-type", "image/png")
	w.Write(img)
}
