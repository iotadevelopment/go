package ixi

type Plugin struct {
    Events pluginEvents
}

func NewPlugin(callback Callback, callbacks ...Callback) *Plugin {
    plugin := &Plugin{
        Events: pluginEvents{
            Configure: &event{make(map[uintptr]Callback)},
            Run:       &event{make(map[uintptr]Callback)},
        },
    }

    if len(callbacks) >= 1 {
        plugin.Events.Configure.Attach(callback)
        for _, callback = range callbacks[:len(callbacks) - 1] {
            plugin.Events.Configure.Attach(callback)
        }

        plugin.Events.Run.Attach(callbacks[len(callbacks) - 1])
    } else {
        plugin.Events.Run.Attach(callback)
    }

    return plugin
}
