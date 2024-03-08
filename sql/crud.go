package sql

type Inserter interface {
	Insert() string
}

type Saver interface {
	Save() string
}

type Deleter interface {
	Delete() string
}

type Updater interface {
	Update() string
}

type Geter interface {
	Get() string
}

type GeterForUpdate interface {
	GetForUpdate() string
}

type Finder interface {
	Find() string
}

type Counter interface {
	Count() string
}

type Pager interface {
	Pagination() string
}

type SQL interface {
	Inserter
	Saver
	Deleter
	Updater
	Geter
	GeterForUpdate
	Finder
	Counter
	Pager
}

type SQLBase struct{}

// Count implements SQL.
func (s *SQLBase) Count() string {
	panic("unimplemented")
}

// Delete implements SQL.
func (s *SQLBase) Delete() string {
	panic("unimplemented")
}

// Find implements SQL.
func (s *SQLBase) Find() string {
	panic("unimplemented")
}

// Get implements SQL.
func (s *SQLBase) Get() string {
	panic("unimplemented")
}

// GetForUpdate implements SQL.
func (s *SQLBase) GetForUpdate() string {
	panic("unimplemented")
}

// Insert implements SQL.
func (s *SQLBase) Insert() string {
	panic("unimplemented")
}

// Pagination implements SQL.
func (s *SQLBase) Pagination() string {
	panic("unimplemented")
}

// Save implements SQL.
func (s *SQLBase) Save() string {
	panic("unimplemented")
}

// Update implements SQL.
func (s *SQLBase) Update() string {
	panic("unimplemented")
}

var _ SQL = (*SQLBase)(nil)
