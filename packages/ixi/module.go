package ixi

type IXIModule struct {
    loadHandlers []LoadHandler
    configureHandlers []ConfigureHandler
    runHandlers []RunHandler
}

func NewIXIModule() *IXIModule {
    return &IXIModule{}
}

func (this *IXIModule) OnLoad(handler LoadHandler) *IXIModule {
    this.loadHandlers = append(this.loadHandlers, handler)

    return this
}

func (this *IXIModule) TriggerLoad() {
    for _, handler := range this.loadHandlers {
        handler()
    }
}

func (this *IXIModule) OnConfigure(handler ConfigureHandler) *IXIModule {
    this.configureHandlers = append(this.configureHandlers, handler)

    return this
}

func (this *IXIModule) TriggerConfigure() {
    for _, handler := range this.configureHandlers {
        handler()
    }
}

func (this *IXIModule) OnRun(handler RunHandler) *IXIModule {
    this.runHandlers = append(this.runHandlers, handler)

    return this
}

func (this *IXIModule) TriggerRun() {
    for _, handler := range this.runHandlers {
        handler()
    }
}