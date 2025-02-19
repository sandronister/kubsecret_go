# KubSecret Generate

KubSecret Generate é uma ferramenta para gerar certificados TLS e arquivos de configuração Kubernetes Secrets a partir de um arquivo PKCS#12 (.p12).

## Instalação

1. Clone o repositório:

```sh
git clone https://github.com/sandronister/kubsecret_generate.git
cd kubsecret_generate
 ```
### Build
Para construir o projeto, execute o seguinte comando:

```sh
go build -o kubsecret_generate
```

#### Uso
1. Coloque seu arquivo PKCS#12 (.p12) no diretório raiz do projeto.
2. Execute o comando:

```sh
./kubsecret_generate <caminho_do_arquivo_p12>
 ```

3. O programa solicitará a senha para abrir o arquivo .p12.
4. Siga as instruções para fornecer o nome e namespace do Secret.