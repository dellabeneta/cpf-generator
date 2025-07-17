package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// CPF representa um CPF brasileiro
type CPF struct {
	Numbers string
	Formatted string
}

// generateCPF gera um CPF válido brasileiro
func generateCPF() CPF {
	rand.Seed(time.Now().UnixNano())
	
	// Gera os primeiros 9 dígitos
	digits := make([]int, 11)
	for i := 0; i < 9; i++ {
		digits[i] = rand.Intn(10)
	}
	
	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	remainder := sum % 11
	if remainder < 2 {
		digits[9] = 0
	} else {
		digits[9] = 11 - remainder
	}
	
	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	remainder = sum % 11
	if remainder < 2 {
		digits[10] = 0
	} else {
		digits[10] = 11 - remainder
	}
	
	// Converte para string
	numbers := ""
	for _, digit := range digits {
		numbers += strconv.Itoa(digit)
	}
	
	// Formata o CPF (XXX.XXX.XXX-XX)
	formatted := fmt.Sprintf("%s.%s.%s-%s", 
		numbers[:3], 
		numbers[3:6], 
		numbers[6:9], 
		numbers[9:11])
	
	return CPF{
		Numbers: numbers,
		Formatted: formatted,
	}
}

// Página HTML
const htmlTemplate = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Gerador de CPF</title>
	<link rel="icon" href="/favicon.ico" type="image/x-icon">

    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 600px;
            margin: 50px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            text-align: center;
            color: #333;
            margin-bottom: 30px;
        }
        .cpf-display {
            background: #e8f4fd;
            padding: 20px;
            border-radius: 8px;
            text-align: center;
            margin: 20px 0;
            border-left: 4px solid #007bff;
        }
        .cpf-number {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
            margin: 10px 0;
            font-family: 'Courier New', monospace;
        }
        .cpf-raw {
            font-size: 14px;
            color: #666;
            font-family: 'Courier New', monospace;
        }
        button {
            background: #007bff;
            color: white;
            border: none;
            padding: 12px 24px;
            font-size: 16px;
            border-radius: 5px;
            cursor: pointer;
            width: 100%;
            margin: 10px 0;
        }
        button:hover {
            background: #0056b3;
        }
        .warning {
            background: #fff3cd;
            color: #856404;
            border: 1px solid #ffeaa7;
            padding: 15px;
            border-radius: 5px;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Gerador de CPF Brasileiro</h1>
        
        <div class="warning">
            <strong>Aviso:</strong> Este gerador é apenas para fins educacionais e de teste. 
            Não use CPFs gerados para fraudes ou atividades ilegais.
        </div>

        {{if .CPF}}
        <div class="cpf-display">
            <div>CPF Gerado:</div>
            <div class="cpf-number">{{.CPF.Formatted}}</div>
            <div class="cpf-raw">Números: {{.CPF.Numbers}}</div>
        </div>
        {{end}}

        <form method="post" action="/">
            <button type="submit" name="action" value="generate">Gerar Novo CPF</button>
        </form>
    </div>
</body>
</html>
`

// PageData contém os dados para renderizar a página
type PageData struct {
	CPF *CPF
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").Parse(htmlTemplate))
	
	if r.Method == "POST" {
		action := r.FormValue("action")
		
		if action == "generate" {
			cpf := generateCPF()
			http.Redirect(w, r, fmt.Sprintf("/?cpf=%s", cpf.Numbers), http.StatusSeeOther)
			return
		}
		
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	data := PageData{}
	
	if len(r.URL.Query()) > 0 {
		if cpfParam := r.URL.Query().Get("cpf"); cpfParam != "" {
			formatted := fmt.Sprintf("%s.%s.%s-%s", 
				cpfParam[:3], 
				cpfParam[3:6], 
				cpfParam[6:9], 
				cpfParam[9:11])
			data.CPF = &CPF{
				Numbers: cpfParam,
				Formatted: formatted,
			}
		}
	}
	
	tmpl.Execute(w, data)
}

func main() {
	http.Handle("/favicon.ico", http.FileServer(http.Dir("static")))
	
	http.HandleFunc("/", homeHandler)
	
	fmt.Println("Servidor iniciado em http://localhost:8080")
	fmt.Println("Pressione Ctrl+C para parar o servidor")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}