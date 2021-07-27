package models

import "github.com/aseabra/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConexaoDB()

	selectProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	insertProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)") //para postgres usa-se $1, $2... ao inves de ?, ?....
	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConexaoDB()

	deletaProduto, err := db.Prepare("delete from produtos where id= ?")
	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(id)

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConexaoDB()

	produtoDoBanco, err := db.Query("select * from produtos where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	AtualizaProduto, err := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
