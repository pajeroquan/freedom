package repositorys

import (
	"github.com/8treenet/freedom"
)

func init() {
	freedom.Prepare(func(initiator freedom.Initiator) {
		initiator.BindRepository(func() *DefaultRepository {
			return &DefaultRepository{}
		})
	})
}

// DefaultRepository .
type DefaultRepository struct {
	freedom.Repository
}

// GetIP .
func (repo *DefaultRepository) GetIP() string {
	//repo.DB().Find()
	repo.Worker.Logger().Infof("我是Repository GetIP")
	return repo.Worker.IrisContext().RemoteAddr()
}

// GetUA - implment DefaultRepoInterface interface
func (repo *DefaultRepository) GetUA() string {
	repo.Worker.Logger().Infof("我是Repository GetUA")
	return repo.Worker.IrisContext().Request().UserAgent()
}
