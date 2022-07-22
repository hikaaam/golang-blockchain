<h1 style="background-color:#90d8f7;color:white;font-family:consolas;text-align:center;padding-top:20px;padding-bottom:20px;">Learning Go</h1>

<h2 style="text-align:center;margin-bottom:30px">Simple BlockChain With GO</h2>
<div style="display:flex;justify-content:center;">
    <img src="https://go.dev/images/gophers/motorcycle.svg" alt="Gopher Logo 1" width="300"/>
</div>

---

### How To Run
build and run the executable
`go build src/main.go`
`main.exe`
or
`go run src/main.go`

---

### API Docs

```javascript
GET
localhost:8080
//will return current running blockchain
```

```typescript
POST
localhost:8080/add
interface Body {
    sender: string;
    receiver: string;
    amount: number; //integer
}
//will add new transaction to the blockchain
```

---

<div style="display:flex;justify-content:center;">
    <img src="https://go.dev/images/gophers/ladder.svg" alt="Gopher Logo 2" width="300"/>
</div>