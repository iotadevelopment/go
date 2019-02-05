package ixi

var loadedPlugins = make([]*Plugin, 0)

func Load(plugins ...*Plugin) {
    loadedPlugins = append(loadedPlugins, plugins...)
}

func Run() {
    for _, plugin := range loadedPlugins {
        plugin.Events.Configure.Trigger()
    }

    for _, plugin := range loadedPlugins {
        plugin.Events.Run.Trigger()
    }
}
