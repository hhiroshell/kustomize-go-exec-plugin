# Build a plugin binary
build:
	go build -o plugin/hhiroshell.github.com/v1/secretgenerator/SecretGenerator plugin/hhiroshell.github.com/v1/secretgenerator/SecretGenerator.go && \
	go build -o plugin/hhiroshell.github.com/v1/lineinsertiontransformer/LineInsertionTransformer plugin/hhiroshell.github.com/v1/lineinsertiontransformer/LineInsertionTransformer.go
