package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"nlpconverter/converter"
	"os"
	"strconv"
	"strings"
)

var testCases = []string{
	// Volume
	"1L & 23 ml",
	"1Liter + 100.87 milli in cm^3",
	".5 gal in L",
	"1.5e3 ml in L",
	"1/2 gallon + 1/4 pint in cups",
	"two pints and a half cup in floz",
	"500ml - .25L",
	"1 leter in ml",
	"2 gallens in L",
	"one gallon and 2.5 litres in ml",

	// Length
	"1 km in miles",
	"a foot and 5 inches in cm",
	"100 meters + 0.1km in ft",

	// Weight
	"1kg in lbs",
	"two pounds and 8 ounces in grams",
	"100g + .5kg",

	// Temperature
	"100 C in F",
	"212 f in C",
	"0c in k",

	// Area
	"100 sqft in m2",
	"2 acres in ha",

	// Speed
	"60 mph in kph",
	"100 km/h in knots",

	// Compound
	"10 km / 2 hr in m/s",

	// Time
	"1 day in hours",

	// Previously failing
	"two pints + a half cup in floz",
	"one gallon + 2.5 litres in ml",
	"a foot + 5 inches in cm",
	"two pounds + 8 ounces in grams",
}

func printHelp() {
	fmt.Println("Usage: nlp-unit-converter [expression]")
	fmt.Println("       nlp-unit-converter [flags]")
	fmt.Println("\nEvaluates a natural language expression of units and converts them.")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help\t\t\tPrints this help message.")
	fmt.Println("  -ss, --start-server [port]\tStarts a web API server (default port: 8080).")
	fmt.Println("\nServer Examples:")
	fmt.Println("  nlp-unit-converter -ss\t\tStart server on default port 8080")
	fmt.Println("  nlp-unit-converter --start-server 7000\tStart server on port 7000")
	fmt.Println("\nConversion Examples:")
	fmt.Println("| Expression                           | Result                                  |")
	fmt.Println("|------------------------------------|-----------------------------------------|")
	unitMap := converter.MustRegisterSystems()
	conv := converter.NewConverter(unitMap)
	for _, tc := range testCases {
		result, err := conv.Process(tc)
		if err != nil {
			fmt.Printf("| %-34s | Error: %-29s |\n", tc, err.Error())
		} else {
			resultStr := fmt.Sprintf("%g %s (%s)", result.Value, result.UnitSymbol, result.UnitName)
			fmt.Printf("| %-34s | %-39s |\n", tc, resultStr)
		}
	}
	fmt.Println("|------------------------------------|-----------------------------------------|")
}

const htmlPage = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>NLP Unit Converter</title><style>*{box-sizing:border-box;margin:0;padding:0}body{font-family:system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;background-color:#f3f4f6;display:flex;align-items:center;justify-content:center;min-height:100vh}.container{width:100%;max-width:448px;margin:1rem;background-color:#fff;border-radius:12px;border:1px solid #e5e7eb;padding:32px}.container>div:not(:first-child){margin-top:24px}h1{font-size:1.5rem;font-weight:700;text-align:center}p{color:#6b7280;text-align:center;margin-top:4px}#expression-input{width:100%;padding:12px 16px;background-color:#f9fafb;border:1px solid #d1d5db;border-radius:8px;font-size:1rem}#expression-input:focus{outline:2px solid #3b82f6}#convert-btn{width:100%;margin-top:16px;background-color:#2563eb;color:#fff;font-weight:600;padding:12px 16px;border:none;border-radius:8px;cursor:pointer}#convert-btn:disabled{background-color:#9ca3af;cursor:not-allowed}#result-display{padding:16px;border-radius:8px;text-align:center;font-weight:500;margin-top:16px}.hidden{display:none}.success{background-color:#d1fae5;color:#065f46}.error{background-color:#fee2e2;color:#991b1b}.examples-section{padding-top:16px;border-top:1px solid #e5e7eb}.examples-section h3{font-size:.875rem;font-weight:600;color:#4b5563;margin-bottom:12px;text-align:center}#examples-list{list-style:none;display:flex;flex-wrap:wrap;justify-content:center;gap:8px}.example-btn{padding:4px 12px;background-color:#f3f4f6;color:#374151;font-size:.875rem;border-radius:9999px;border:1px solid #d1d5db;cursor:pointer}</style></head><body><div class="container"><div><h1>Unit Converter</h1><p>Convert units using natural language.</p></div><div><input type="text" id="expression-input" placeholder="e.g., 2 liters to ml"><button id="convert-btn">Convert</button></div><div id="result-display" class="hidden"></div><div class="examples-section"><h3>Try these:</h3><ul id="examples-list"></ul></div></div><script>const expressionInput = document.getElementById('expression-input');
        const convertBtn = document.getElementById('convert-btn');
        const resultDisplay = document.getElementById('result-display');
        const examplesList = document.getElementById('examples-list');
        const examples = [
            '2l to ml', '100 f to c', '5 km to miles',
            '1 gallon to ml', '100 lbs to kg', '1 day in hours'
        ];
        const handleConversion = async () => {
            const expression = expressionInput.value.trim();
            if (!expression) {
                showResult('Enter something to convert.', false);
                return;
            }
            convertBtn.disabled = true;
            convertBtn.textContent = 'Converting...';
            resultDisplay.classList.add('hidden');
            try {
                const response = await fetch('/?q=' + encodeURIComponent(expression));
                if (!response.ok) {
                    throw new Error('API Error: ' + response.statusText);
                }
                const data = await response.json();
                if (data.error) {
                    showResult('Error: ' + data.error, false);
                } else {
                    const resultText = data.value + ' ' + data.unit_symbol + ' (' + data.unit_name + ')';
                    showResult(resultText, true);
                }
            } catch (error) {
                console.error("Conversion failed:", error);
                showResult('Failed to convert. Check the format or try again.', false);
            } finally {
                convertBtn.disabled = false;
                convertBtn.textContent = 'Convert';
            }
        };
        const showResult = (message, isSuccess) => {
            resultDisplay.textContent = message;
            resultDisplay.className = 'result-display';
            if (isSuccess) {
                resultDisplay.classList.add('success');
            } else {
                resultDisplay.classList.add('error');
            }
        };
        const populateExamples = () => {
            examples.forEach(text => {
                const li = document.createElement('li');
                const button = document.createElement('button');
                button.textContent = text;
                button.className = 'example-btn';
                button.onclick = () => {
                    expressionInput.value = text;
                    expressionInput.focus();
                    handleConversion();
                };
                li.appendChild(button);
                examplesList.appendChild(li);
            });
        };
        convertBtn.addEventListener('click', handleConversion);
        expressionInput.addEventListener('keypress', (event) => {
            if (event.key === 'Enter') {
                handleConversion();
            }
        });
        document.addEventListener('DOMContentLoaded', populateExamples);</script></body></html>`

type APIResponse struct {
	Value      float64 `json:"value"`
	UnitSymbol string  `json:"unit_symbol"`
	UnitName   string  `json:"unit_name"`
	Error      string  `json:"error,omitempty"`
}

func startServer(port int) {
	unitMap := converter.MustRegisterSystems()
	conv := converter.NewConverter(unitMap)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(APIResponse{Error: "Only GET requests are allowed"})
			return
		}

		query := r.URL.Query().Get("q")

		// If no query parameter, return HTML page
		if query == "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(htmlPage))
			return
		}

		// Limit query to 100 characters
		if len(query) > 100 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Error: "Query too long. Maximum 100 characters allowed."})
			return
		}

		// Process the conversion
		result, err := conv.Process(query)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Error: err.Error()})
		} else {
			json.NewEncoder(w).Encode(APIResponse{
				Value:      result.Value,
				UnitSymbol: result.UnitSymbol,
				UnitName:   result.UnitName,
			})
		}
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("🚀 NLP Unit Converter API Server starting on http://localhost:%d\n", port)
	log.Printf("📝 Open http://localhost:%d in your browser\n", port)
	log.Printf("🔧 API endpoint: http://localhost:%d/?q=your+query\n", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server on port %d: %v\n", port, err)
	}
}

func main() {
	help := flag.Bool("h", false, "Prints the help message.")
	flag.BoolVar(help, "help", false, "Prints the help message.")
	serverMode := flag.Bool("ss", false, "Starts a web API server.")
	flag.BoolVar(serverMode, "start-server", false, "Starts a web API server.")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	if *serverMode {
		port := 8080

		args := flag.Args()
		if len(args) > 0 {
			if portNum, err := strconv.Atoi(args[0]); err == nil && portNum > 0 && portNum <= 65535 {
				port = portNum
			} else {
				log.Fatalf("❌ Invalid port number: %s. Port must be between 1 and 65535.\n", args[0])
			}
		}

		startServer(port)
		return
	}

	unitMap := converter.MustRegisterSystems()
	conv := converter.NewConverter(unitMap)

	if len(os.Args) == 1 {
		fmt.Println("No expression provided. Use -h or --help for usage information.")
		os.Exit(1)
	}

	expression := strings.Join(flag.Args(), " ")
	if expression == "" {
		fmt.Println("No expression provided. Use -h or --help for usage information.")
		os.Exit(1)
	}

	result, err := conv.Process(expression)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%g %s (%s)\n", result.Value, result.UnitSymbol, result.UnitName)
		os.Exit(0)
	}
}
