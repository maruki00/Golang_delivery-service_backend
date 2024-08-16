package app

type App struct {
	objs map[string]any
}

// func (app *App) Get(key interface{}) any {
// 	obj, ok := app.objs[key.(any)]
// 	if !ok {
// 		app.objs[key] = new(key)
// 	}
// }
