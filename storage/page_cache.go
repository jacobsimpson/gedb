package storage

type PageId int

type Page interface {
	Get(rowid RowId) Row
	Scan() RowScanner
	NextPage() PageId
}

type page struct{}

func readPage(pageId PageId) Page {
	return nil
}

func (p *page) Get(rowid RowId) Row { return nil }

func (p *page) Scan() RowScanner { return nil }
