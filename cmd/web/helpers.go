package main

func (app *application) background(fn func()) {
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.errorLogger.Printf("%v", err)
			}
		}()

		fn()
	}()
}
