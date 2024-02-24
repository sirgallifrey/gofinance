// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package shell

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "gofinance/web/view"

func hotReloadScript() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_hotReloadScript_f5fe`,
		Function: `function __templ_hotReloadScript_f5fe(){(function () {
  var lastUuid = "";
  var timeout;

  const resetBackoff = () => {
    timeout = 1000;
  };

  const backOff = () => {
    if (timeout > 10 * 1000) {
      return;
    }

    timeout = timeout * 2;
  };

  const hotReloadUrl = () => {
    const hostAndPort =
      location.hostname + (location.port ? ":" + location.port : "");

    if (location.protocol === "https:") {
      return "wss://" + hostAndPort + "/ws/hotreload";
    } else if (location.protocol === "http:") {
      return "ws://" + hostAndPort + "/ws/hotreload";
    }
  };

  function connectHotReload() {
    const socket = new WebSocket(hotReloadUrl());

    socket.onmessage = (event) => {
      if (lastUuid === "") {
        lastUuid = event.data;
      }

      if (lastUuid !== event.data) {
        console.log("[Hot Reloader] Server Changed, reloading");
        location.reload();
      }
    };

    socket.onopen = () => {
      resetBackoff();
      socket.send("Hello");
    };

    socket.onclose = () => {
      const timeoutId = setTimeout(function () {
        clearTimeout(timeoutId);
        backOff();

        connectHotReload();
      }, timeout);
    };
  }

  resetBackoff();
  connectHotReload();
})();}`,
		Call:       templ.SafeScript(`__templ_hotReloadScript_f5fe`),
		CallInline: templ.SafeScriptInline(`__templ_hotReloadScript_f5fe`),
	}
}

func HotReload() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if view.IsLocal(ctx) {
			templ_7745c5c3_Err = hotReloadScript().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
