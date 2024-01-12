package meetingus

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-go-modules/modules/meetingus/const4meetingus"
	"net/http"
)

func Module() modules.Module {
	return modules.NewModule(const4meetingus.ModuleID, func(handle modules.HTTPHandleFunc) {
		handle("POST", "/api4meetingus/about", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte("meetingus"))
		})
	})
}
