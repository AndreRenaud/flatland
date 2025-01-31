package editor

import (
	"reflect"

	"github.com/bradbev/flatland/src/asset"
	"github.com/bradbev/flatland/src/editor/edgui"
	"github.com/bradbev/flatland/src/flat"

	"github.com/inkyblackness/imgui-go/v4"
)

type assetEditWindow struct {
	target  asset.Asset
	path    string
	context *TypeEditContext
}

func newAssetEditWindow(path string, target asset.Asset, context *TypeEditContext) *assetEditWindow {
	if playable, ok := target.(flat.Playable); ok {
		playable.BeginPlay()
	}
	return &assetEditWindow{
		path:    path,
		target:  target,
		context: context,
	}
}

func (a *assetEditWindow) Draw() error {
	defer imgui.End()
	open := true
	if imgui.BeginV(a.path, &open, 0) {
		enabled := a.context.hasChanged
		edgui.WithDisabled(!enabled, func() {
			imgui.SameLineV(0, imgui.WindowWidth()-180)
			reload := false
			if imgui.Button("Save") && enabled {
				asset.Save(asset.Path(a.path), a.target)
				a.context.hasChanged = false
				reload = true
			}
			imgui.SameLine()
			if (imgui.Button("Revert") || reload) && enabled {
				asset.LoadWithOptions(asset.Path(a.path), asset.LoadOptions{ForceReload: true})
				if playable, ok := a.target.(flat.Playable); ok {
					playable.BeginPlay()
				}
			}
		})
		imgui.SameLine()
		if imgui.Button("Refresh") {
			if playable, ok := a.target.(flat.Playable); ok {
				playable.BeginPlay()
			}
		}
		value := reflect.ValueOf(a.target)
		a.context.EditValue(value)
	}
	if !open {
		// If there is context from the editor, delete it
		DisposeContext(a.context, reflect.ValueOf(a.target))
		return closeDrawable
	}
	return nil
}
