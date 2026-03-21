# Pokedex CLI

Uma Pokedex interativa via terminal, escrita em Go. Explore locais, encontre Pokemon, capture-os e monte sua colecao — tudo direto do seu terminal.

Projeto desenvolvido como parte do curso [Boot.dev](https://boot.dev).

## Funcionalidades

- **Navegacao por locais** — percorra as areas do mundo Pokemon com paginacao (`map` / `mapb`)
- **Exploracao** — veja quais Pokemon aparecem em cada area
- **Captura** — tente capturar Pokemon com mecanica probabilistica baseada na experiencia base
- **Inspecao** — consulte stats, tipos e detalhes dos Pokemon capturados
- **Pokedex** — visualize todos os Pokemon que voce ja capturou
- **Cache inteligente** — respostas da API sao cacheadas em memoria com TTL de 5 minutos

## Requisitos

- Go 1.26+

## Instalacao e Execucao

```bash
git clone https://github.com/viniciusLambert/Pokedex.git
cd Pokedex
go build -o pokedex && ./pokedex
```

Ou simplesmente:

```bash
go run .
```

## Comandos

| Comando | Descricao |
|---------|-----------|
| `help` | Exibe todos os comandos disponiveis |
| `map` | Lista as proximas 20 areas de localizacao |
| `mapb` | Lista as 20 areas anteriores |
| `explore <area>` | Lista os Pokemon encontrados em uma area |
| `catch <pokemon>` | Tenta capturar um Pokemon |
| `inspect <pokemon>` | Exibe detalhes de um Pokemon capturado |
| `pokedex` | Mostra todos os Pokemon capturados |
| `exit` | Sai do programa |

## Exemplo de Uso

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
 - tentacool
 - tentacruel
 - magikarp
 - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  - hp: 20
  - attack: 10
  - defense: 55
Types:
  - water

Pokedex > pokedex
 - magikarp
```

## Arquitetura

```
.
├── main.go                    # Ponto de entrada
├── repl.go                    # REPL e roteamento de comandos
├── command_*.go               # Implementacao dos comandos
└── internal/
    ├── pokeapi/               # Client HTTP para a PokeAPI
    │   ├── client.go          # Client com cache embutido
    │   ├── fetchLocationList.go
    │   ├── getLocationAreas.go
    │   ├── getPokemon.go
    │   └── types_*.go         # Structs de dados
    └── pokecache/             # Cache em memoria com TTL
        ├── cache.go
        ├── add.go
        ├── get.go
        └── reaploop.go        # Goroutine de expiracao
```

## Proximos Passos

- [ ] Suporte a setas (historico de comandos)
- [ ] Batalhas entre Pokemon
- [ ] Mais testes unitarios
- [ ] Sistema de party com level up
- [ ] Evolucao de Pokemon
- [ ] Persistencia da Pokedex em disco (salvar progresso entre sessoes)
- [ ] Navegacao interativa de areas (escolher direcoes em vez de digitar nomes)
- [ ] Encontros aleatorios com Pokemon selvagens
- [ ] Diferentes tipos de Pokebolas (Great Ball, Ultra Ball) com taxas de captura variadas

## Tecnologias

- **Go** — linguagem principal
- **[PokeAPI](https://pokeapi.co)** — API publica de dados Pokemon
- **Cache custom** — implementacao propria com TTL e goroutines, thread-safe com `sync.Mutex`
