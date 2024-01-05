# AutoPrinter

## Português / Portuguese

Este sistema foi desenvolvido para automatizar o processo de impressão de imagens em uma impressora específica, registrando informações relevantes em um arquivo de log. O sistema possui dois scripts, um em Bash e outro em PowerShell, para sistemas Unix-like e Windows, respectivamente.
Funcionalidades:

- Automatização da Impressão:
        Os scripts aceitam o nome da imagem como parâmetro e verificam se o arquivo existe no diretório especificado.
        Caso a extensão da imagem não seja fornecida, o sistema acrescenta automaticamente a extensão '.jpeg' para facilitar o processo.

- Registro de Logs:
        Durante o processo de impressão, os scripts registram as seguintes informações no arquivo de log:
            Data e hora de início da impressão.
            Informações sobre a imagem e a impressora utilizada.
            Data e hora de conclusão da impressão.

### Uso:

- Bash (Sistemas Unix-like):
        Para utilizar o script em Bash, execute-o no terminal fornecendo o nome do arquivo como argumento. Exemplo:

```bash
./autoprinter.sh "nome_da_imagem"
```

- PowerShell (Windows):

    No PowerShell, execute o script fornecendo o nome do arquivo como argumento. Exemplo:

```arduino
.\autoprinter.ps1 "nome_da_imagem"
```

### API Server (Go):

Para rodar a aplicação de forma mais eficiente e profissional, vamos usar uma API em Go para automatizar a execução dos scripts.

Você pode usar o comando `go run main.go` para rodar a aplicação, ou `go build main.go` para gerar um executável.

#### Build

Caso você queira buildar seu projeto, rode o seguinte comando no terminal:

```bash
go build -o <nome_do_executavel>
```

No nosso caso, estamos usando o nome `autoprinter`, então o comando fica:

```bash
go build -o autoprinter
```

Lembre-se de que após o build, o arquivo só pode ser executado em sistemas operacionais compatíveis com o seu sistema de desenvolvimento. Por exemplo, se você estiver desenvolvendo no Windows, o executável só poderá ser executado em sistemas Windows.

Mas caso você queira buildar para outros sistemas, você pode usar o seguinte comando:

```bash
GOOS=<sistema_operacional> GOARCH=<arquitetura> go build -o <nome_do_executavel>
```

Por exemplo, para buildar para Windows, você pode usar:

```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o nome_do_arquivo_windows.exe
```

Caso você queira buildar para Linux, você pode usar:

```bash
GOOS=linux GOARCH=amd64 go build -o nome_do_arquivo_linux
```

Muitas vezes, o comando acima não funciona, por que o compilador não está instalado. Para instalar o compilador, você pode usar o seguinte comando:

```bash
sudo apt install gcc-mingw-w64-x86-64
```

#### Execução

Para executar o programa, você pode usar o seguinte comando:

```bash
./autoprinter
```

Caso você tenha buildado para outro sistema, você pode usar o seguinte comando:

```bash
./nome_do_executavel
```

Logo após a execução, o programa irá rodar na porta 8080. Para acessar a API, você pode usar o seguinte link:

```bash
http://localhost:8080/print
```

Lá ira conter as instruções para a utilização da API.

#### Observações

Pode ser que ao rodar no Windows, você receba um erro de permissão. Para resolver isso, você pode usar o seguinte comando:

```bash
Set-ExecutionPolicy RemoteSigned
```

---

## Inglês / English

This system was developed to automate the process of printing images on a specific printer, logging relevant information to a log file. The system consists of two scripts, one in Bash for Unix-like systems and another in PowerShell for Windows.
Features:

- Automated Printing:
        The scripts accept the image name as a parameter and check if the file exists in the specified directory.
        If the image extension is not provided, the system automatically appends the '.jpeg' extension for convenience.

- Logging:
        During the printing process, the scripts log the following information to the log file:
            Start date and time of the printing process.
            Information about the image and the printer used.
            Completion date and time of the printing process.

### Usage:

- Bash (Unix-like Systems):
        To use the Bash script, execute it in the terminal providing the image name as an argument. Example:

```bash
./autoprinter.sh "image_name"
```

- PowerShell (Windows):

    In PowerShell, execute the script providing the image name as an argument. Example:

```arduino
.\autoprinter.ps1 "image_name"
```

### API Server (Go):

You can use the command go run main.go to run the application, or go build main.go to generate an executable.
Build

If you want to build your project, run the following command in the terminal:

```bash
go build -o <executable_name>
```

In our case, we're using the name autoprinter, so the command is:

```bash
go build -o autoprinter
```

Remember that after the build, the file can only be executed on operating systems compatible with your development system. For example, if you're developing on Windows, the executable can only be run on Windows systems.

But if you want to build for other systems, you can use the following command:

```bash
GOOS=<operating_system> GOARCH=<architecture> go build -o <executable_name>
```

For example, to build for Windows, you can use:

```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o filename_windows.exe
```

If you want to build for Linux, you can use:

```bash
GOOS=linux GOARCH=amd64 go build -o filename_linux
```

Often, the above command doesn't work because the compiler isn't installed. To install the compiler, you can use the following command:

```bash
sudo apt install gcc-mingw-w64-x86-64
```

### Execution

To run the program, you can use the following command:

```bash
./autoprinter
```

If you've built for another system, you can use the following command:

```bash
./executable_name
```

Immediately after execution, the program will run on port 8080. To access the API, you can use the following link:

```bash
http://localhost:8080/print
```

There, you'll find instructions for using the API.

### Notes

When running on Windows, you may receive a permission error. To fix this, you can use the following command:

```bash
Set-ExecutionPolicy RemoteSigned
```
