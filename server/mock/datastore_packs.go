// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import "github.com/kolide/fleet/server/kolide"

var _ kolide.PackStore = (*PackStore)(nil)

type ApplyPackSpecsFunc func(specs []*kolide.PackSpec) error

type GetPackSpecsFunc func() ([]*kolide.PackSpec, error)

type GetPackSpecFunc func(name string) (*kolide.PackSpec, error)

type DeletePackFunc func(name string) error

type PackFunc func(pid uint) (*kolide.Pack, error)

type ListPacksFunc func(opt kolide.ListOptions) ([]*kolide.Pack, error)

type PackByNameFunc func(name string, opts ...kolide.OptionalArg) (*kolide.Pack, bool, error)

type ListLabelsForPackFunc func(pid uint) ([]*kolide.Label, error)

type ListPacksForHostFunc func(hid uint) (packs []*kolide.Pack, err error)

type ListHostsInPackFunc func(pid uint, opt kolide.ListOptions) ([]uint, error)

type ListExplicitHostsInPackFunc func(pid uint, opt kolide.ListOptions) ([]uint, error)

type PackStore struct {
	ApplyPackSpecsFunc        ApplyPackSpecsFunc
	ApplyPackSpecsFuncInvoked bool

	GetPackSpecsFunc        GetPackSpecsFunc
	GetPackSpecsFuncInvoked bool

	GetPackSpecFunc        GetPackSpecFunc
	GetPackSpecFuncInvoked bool

	DeletePackFunc        DeletePackFunc
	DeletePackFuncInvoked bool

	PackFunc        PackFunc
	PackFuncInvoked bool

	ListPacksFunc        ListPacksFunc
	ListPacksFuncInvoked bool

	PackByNameFunc        PackByNameFunc
	PackByNameFuncInvoked bool

	ListLabelsForPackFunc        ListLabelsForPackFunc
	ListLabelsForPackFuncInvoked bool

	ListPacksForHostFunc        ListPacksForHostFunc
	ListPacksForHostFuncInvoked bool

	ListHostsInPackFunc        ListHostsInPackFunc
	ListHostsInPackFuncInvoked bool

	ListExplicitHostsInPackFunc        ListExplicitHostsInPackFunc
	ListExplicitHostsInPackFuncInvoked bool
}

func (s *PackStore) ApplyPackSpecs(specs []*kolide.PackSpec) error {
	s.ApplyPackSpecsFuncInvoked = true
	return s.ApplyPackSpecsFunc(specs)
}

func (s *PackStore) GetPackSpecs() ([]*kolide.PackSpec, error) {
	s.GetPackSpecsFuncInvoked = true
	return s.GetPackSpecsFunc()
}

func (s *PackStore) GetPackSpec(name string) (*kolide.PackSpec, error) {
	s.GetPackSpecFuncInvoked = true
	return s.GetPackSpecFunc(name)
}

func (s *PackStore) DeletePack(name string) error {
	s.DeletePackFuncInvoked = true
	return s.DeletePackFunc(name)
}

func (s *PackStore) Pack(pid uint) (*kolide.Pack, error) {
	s.PackFuncInvoked = true
	return s.PackFunc(pid)
}

func (s *PackStore) ListPacks(opt kolide.ListOptions) ([]*kolide.Pack, error) {
	s.ListPacksFuncInvoked = true
	return s.ListPacksFunc(opt)
}

func (s *PackStore) PackByName(name string, opts ...kolide.OptionalArg) (*kolide.Pack, bool, error) {
	s.PackByNameFuncInvoked = true
	return s.PackByNameFunc(name, opts...)
}

func (s *PackStore) ListLabelsForPack(pid uint) ([]*kolide.Label, error) {
	s.ListLabelsForPackFuncInvoked = true
	return s.ListLabelsForPackFunc(pid)
}

func (s *PackStore) ListPacksForHost(hid uint) (packs []*kolide.Pack, err error) {
	s.ListPacksForHostFuncInvoked = true
	return s.ListPacksForHostFunc(hid)
}

func (s *PackStore) ListHostsInPack(pid uint, opt kolide.ListOptions) ([]uint, error) {
	s.ListHostsInPackFuncInvoked = true
	return s.ListHostsInPackFunc(pid, opt)
}

func (s *PackStore) ListExplicitHostsInPack(pid uint, opt kolide.ListOptions) ([]uint, error) {
	s.ListExplicitHostsInPackFuncInvoked = true
	return s.ListExplicitHostsInPackFunc(pid, opt)
}
