# AutoPrinter

### Português / Portuguese

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

### Inglês / English

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