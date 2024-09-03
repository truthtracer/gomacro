// this file was generated by gomacro command: import _i "github.com/go-rod/rod"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package thirdparty

import (
	context "context"
	rod_ "github.com/go-rod/rod"
	cdp "github.com/go-rod/rod/lib/cdp"
	r "reflect"
)

/**
 * Declare and fill a global variable Packages, whose type is compatible
 * with the global variable github.com/truthtracer/gomacro/imports.Packages
 *
 * If you want to automatically register this package's declarations into
 *   github.com/truthtracer/gomacro/imports.Packages
 * to let gomacro know about this package, and allow importing it without compiling
 * a plugin, you can add the following to some _other_ file in this directory:
 *
 * import "github.com/truthtracer/gomacro/imports"
 *
 * func init() {
 *     for k, v := range Packages {
 *         imports.Packages[k] = v
 *     }
 * }
 *
 * Such code is _not_ automatically added to this file, because it would introduce
 * a dependency on gomacro packages, which may be undesiderable.
 */

// reflection: allow interpreted code to import "github.com/go-rod/rod"
func init() {
	Packages["github.com/go-rod/rod"] = Package{
		Name: "rod",
		Binds: map[string]r.Value{
			"DefaultLogger":             r.ValueOf(&rod_.DefaultLogger).Elem(),
			"DefaultSleeper":            r.ValueOf(&rod_.DefaultSleeper).Elem(),
			"Eval":                      r.ValueOf(rod_.Eval),
			"KeyActionPress":            r.ValueOf(rod_.KeyActionPress),
			"KeyActionRelease":          r.ValueOf(rod_.KeyActionRelease),
			"KeyActionTypeKey":          r.ValueOf(rod_.KeyActionTypeKey),
			"New":                       r.ValueOf(rod_.New),
			"NewBrowserPool":            r.ValueOf(rod_.NewBrowserPool),
			"NewPagePool":               r.ValueOf(rod_.NewPagePool),
			"NewStreamReader":           r.ValueOf(rod_.NewStreamReader),
			"NotFoundSleeper":           r.ValueOf(rod_.NotFoundSleeper),
			"SelectorTypeCSSSector":     r.ValueOf(rod_.SelectorTypeCSSSector),
			"SelectorTypeRegex":         r.ValueOf(rod_.SelectorTypeRegex),
			"SelectorTypeText":          r.ValueOf(rod_.SelectorTypeText),
			"TraceTypeInput":            r.ValueOf(rod_.TraceTypeInput),
			"TraceTypeQuery":            r.ValueOf(rod_.TraceTypeQuery),
			"TraceTypeWait":             r.ValueOf(rod_.TraceTypeWait),
			"TraceTypeWaitRequests":     r.ValueOf(rod_.TraceTypeWaitRequests),
			"TraceTypeWaitRequestsIdle": r.ValueOf(rod_.TraceTypeWaitRequestsIdle),
			"Try":                       r.ValueOf(rod_.Try),
		}, Types: map[string]r.Type{
			"Browser":                 r.TypeOf((*rod_.Browser)(nil)).Elem(),
			"CDPClient":               r.TypeOf((*rod_.CDPClient)(nil)).Elem(),
			"CoveredError":            r.TypeOf((*rod_.CoveredError)(nil)).Elem(),
			"Element":                 r.TypeOf((*rod_.Element)(nil)).Elem(),
			"ElementNotFoundError":    r.TypeOf((*rod_.ElementNotFoundError)(nil)).Elem(),
			"Elements":                r.TypeOf((*rod_.Elements)(nil)).Elem(),
			"EvalError":               r.TypeOf((*rod_.EvalError)(nil)).Elem(),
			"EvalOptions":             r.TypeOf((*rod_.EvalOptions)(nil)).Elem(),
			"ExpectElementError":      r.TypeOf((*rod_.ExpectElementError)(nil)).Elem(),
			"ExpectElementsError":     r.TypeOf((*rod_.ExpectElementsError)(nil)).Elem(),
			"Hijack":                  r.TypeOf((*rod_.Hijack)(nil)).Elem(),
			"HijackRequest":           r.TypeOf((*rod_.HijackRequest)(nil)).Elem(),
			"HijackResponse":          r.TypeOf((*rod_.HijackResponse)(nil)).Elem(),
			"HijackRouter":            r.TypeOf((*rod_.HijackRouter)(nil)).Elem(),
			"InvisibleShapeError":     r.TypeOf((*rod_.InvisibleShapeError)(nil)).Elem(),
			"KeyAction":               r.TypeOf((*rod_.KeyAction)(nil)).Elem(),
			"KeyActionType":           r.TypeOf((*rod_.KeyActionType)(nil)).Elem(),
			"KeyActions":              r.TypeOf((*rod_.KeyActions)(nil)).Elem(),
			"Keyboard":                r.TypeOf((*rod_.Keyboard)(nil)).Elem(),
			"Message":                 r.TypeOf((*rod_.Message)(nil)).Elem(),
			"Mouse":                   r.TypeOf((*rod_.Mouse)(nil)).Elem(),
			"NavigationError":         r.TypeOf((*rod_.NavigationError)(nil)).Elem(),
			"NoPointerEventsError":    r.TypeOf((*rod_.NoPointerEventsError)(nil)).Elem(),
			"NoShadowRootError":       r.TypeOf((*rod_.NoShadowRootError)(nil)).Elem(),
			"NotInteractableError":    r.TypeOf((*rod_.NotInteractableError)(nil)).Elem(),
			"ObjectNotFoundError":     r.TypeOf((*rod_.ObjectNotFoundError)(nil)).Elem(),
			"Page":                    r.TypeOf((*rod_.Page)(nil)).Elem(),
			"PageCloseCanceledError":  r.TypeOf((*rod_.PageCloseCanceledError)(nil)).Elem(),
			"PageNotFoundError":       r.TypeOf((*rod_.PageNotFoundError)(nil)).Elem(),
			"Pages":                   r.TypeOf((*rod_.Pages)(nil)).Elem(),
			"RaceContext":             r.TypeOf((*rod_.RaceContext)(nil)).Elem(),
			"ScrollScreenshotOptions": r.TypeOf((*rod_.ScrollScreenshotOptions)(nil)).Elem(),
			"SearchResult":            r.TypeOf((*rod_.SearchResult)(nil)).Elem(),
			"SelectorType":            r.TypeOf((*rod_.SelectorType)(nil)).Elem(),
			"StreamReader":            r.TypeOf((*rod_.StreamReader)(nil)).Elem(),
			"Touch":                   r.TypeOf((*rod_.Touch)(nil)).Elem(),
			"TraceType":               r.TypeOf((*rod_.TraceType)(nil)).Elem(),
			"TryError":                r.TypeOf((*rod_.TryError)(nil)).Elem(),
		}, Proxies: map[string]r.Type{
			"CDPClient": r.TypeOf((*P__root_go_src_rod_CDPClient)(nil)).Elem(),
		}, Wrappers: map[string][]string{
			"CoveredError":         []string{"Attribute", "BackgroundImage", "Blur", "Call", "CancelTimeout", "CanvasToImage", "Click", "ContainsElement", "Context", "Describe", "Disabled", "Element", "ElementByJS", "ElementR", "ElementX", "Elements", "ElementsByJS", "ElementsX", "Equal", "Eval", "Evaluate", "Focus", "Frame", "GetContext", "GetSessionID", "GetXPath", "HTML", "Has", "HasR", "HasX", "Hover", "Input", "InputColor", "InputTime", "Interactable", "KeyActions", "Matches", "MoveMouseOut", "MustAttribute", "MustBackgroundImage", "MustBlur", "MustCanvasToImage", "MustClick", "MustContainsElement", "MustDescribe", "MustDisabled", "MustDoubleClick", "MustElement", "MustElementByJS", "MustElementR", "MustElementX", "MustElements", "MustElementsByJS", "MustElementsX", "MustEqual", "MustEval", "MustFocus", "MustFrame", "MustGetXPath", "MustHTML", "MustHas", "MustHasR", "MustHasX", "MustHover", "MustInput", "MustInputColor", "MustInputTime", "MustInteractable", "MustKeyActions", "MustMatches", "MustMoveMouseOut", "MustNext", "MustParent", "MustParents", "MustPrevious", "MustProperty", "MustRelease", "MustRemove", "MustResource", "MustScreenshot", "MustScrollIntoView", "MustSelect", "MustSelectAllText", "MustSelectText", "MustSetFiles", "MustShadowRoot", "MustShape", "MustTap", "MustText", "MustType", "MustVisible", "MustWait", "MustWaitEnabled", "MustWaitInteractable", "MustWaitInvisible", "MustWaitLoad", "MustWaitStable", "MustWaitVisible", "MustWaitWritable", "Next", "Overlay", "Page", "Parent", "Parents", "Previous", "Property", "Release", "Remove", "Resource", "Screenshot", "ScrollIntoView", "Select", "SelectAllText", "SelectText", "SetFiles", "ShadowRoot", "Shape", "Sleeper", "String", "Tap", "Text", "Timeout", "Type", "Visible", "Wait", "WaitEnabled", "WaitInteractable", "WaitInvisible", "WaitLoad", "WaitStable", "WaitStableRAF", "WaitVisible", "WaitWritable", "WithCancel", "WithPanic"},
			"InvisibleShapeError":  []string{"Attribute", "BackgroundImage", "Blur", "Call", "CancelTimeout", "CanvasToImage", "Click", "ContainsElement", "Context", "Describe", "Disabled", "Element", "ElementByJS", "ElementR", "ElementX", "Elements", "ElementsByJS", "ElementsX", "Equal", "Eval", "Evaluate", "Focus", "Frame", "GetContext", "GetSessionID", "GetXPath", "HTML", "Has", "HasR", "HasX", "Hover", "Input", "InputColor", "InputTime", "Interactable", "KeyActions", "Matches", "MoveMouseOut", "MustAttribute", "MustBackgroundImage", "MustBlur", "MustCanvasToImage", "MustClick", "MustContainsElement", "MustDescribe", "MustDisabled", "MustDoubleClick", "MustElement", "MustElementByJS", "MustElementR", "MustElementX", "MustElements", "MustElementsByJS", "MustElementsX", "MustEqual", "MustEval", "MustFocus", "MustFrame", "MustGetXPath", "MustHTML", "MustHas", "MustHasR", "MustHasX", "MustHover", "MustInput", "MustInputColor", "MustInputTime", "MustInteractable", "MustKeyActions", "MustMatches", "MustMoveMouseOut", "MustNext", "MustParent", "MustParents", "MustPrevious", "MustProperty", "MustRelease", "MustRemove", "MustResource", "MustScreenshot", "MustScrollIntoView", "MustSelect", "MustSelectAllText", "MustSelectText", "MustSetFiles", "MustShadowRoot", "MustShape", "MustTap", "MustText", "MustType", "MustVisible", "MustWait", "MustWaitEnabled", "MustWaitInteractable", "MustWaitInvisible", "MustWaitLoad", "MustWaitStable", "MustWaitVisible", "MustWaitWritable", "Next", "Overlay", "Page", "Parent", "Parents", "Previous", "Property", "Release", "Remove", "Resource", "Screenshot", "ScrollIntoView", "Select", "SelectAllText", "SelectText", "SetFiles", "ShadowRoot", "Shape", "Sleeper", "String", "Tap", "Text", "Timeout", "Type", "Visible", "Wait", "WaitEnabled", "WaitInteractable", "WaitInvisible", "WaitLoad", "WaitStable", "WaitStableRAF", "WaitVisible", "WaitWritable", "WithCancel", "WithPanic"},
			"Keyboard":             []string{"Lock", "TryLock", "Unlock"},
			"Mouse":                []string{"Lock", "TryLock", "Unlock"},
			"NoPointerEventsError": []string{"Attribute", "BackgroundImage", "Blur", "Call", "CancelTimeout", "CanvasToImage", "Click", "ContainsElement", "Context", "Describe", "Disabled", "Element", "ElementByJS", "ElementR", "ElementX", "Elements", "ElementsByJS", "ElementsX", "Equal", "Eval", "Evaluate", "Focus", "Frame", "GetContext", "GetSessionID", "GetXPath", "HTML", "Has", "HasR", "HasX", "Hover", "Input", "InputColor", "InputTime", "Interactable", "KeyActions", "Matches", "MoveMouseOut", "MustAttribute", "MustBackgroundImage", "MustBlur", "MustCanvasToImage", "MustClick", "MustContainsElement", "MustDescribe", "MustDisabled", "MustDoubleClick", "MustElement", "MustElementByJS", "MustElementR", "MustElementX", "MustElements", "MustElementsByJS", "MustElementsX", "MustEqual", "MustEval", "MustFocus", "MustFrame", "MustGetXPath", "MustHTML", "MustHas", "MustHasR", "MustHasX", "MustHover", "MustInput", "MustInputColor", "MustInputTime", "MustInteractable", "MustKeyActions", "MustMatches", "MustMoveMouseOut", "MustNext", "MustParent", "MustParents", "MustPrevious", "MustProperty", "MustRelease", "MustRemove", "MustResource", "MustScreenshot", "MustScrollIntoView", "MustSelect", "MustSelectAllText", "MustSelectText", "MustSetFiles", "MustShadowRoot", "MustShape", "MustTap", "MustText", "MustType", "MustVisible", "MustWait", "MustWaitEnabled", "MustWaitInteractable", "MustWaitInvisible", "MustWaitLoad", "MustWaitStable", "MustWaitVisible", "MustWaitWritable", "Next", "Overlay", "Page", "Parent", "Parents", "Previous", "Property", "Release", "Remove", "Resource", "Screenshot", "ScrollIntoView", "Select", "SelectAllText", "SelectText", "SetFiles", "ShadowRoot", "Shape", "Sleeper", "String", "Tap", "Text", "Timeout", "Type", "Visible", "Wait", "WaitEnabled", "WaitInteractable", "WaitInvisible", "WaitLoad", "WaitStable", "WaitStableRAF", "WaitVisible", "WaitWritable", "WithCancel", "WithPanic"},
			"NoShadowRootError":    []string{"Attribute", "BackgroundImage", "Blur", "Call", "CancelTimeout", "CanvasToImage", "Click", "ContainsElement", "Context", "Describe", "Disabled", "Element", "ElementByJS", "ElementR", "ElementX", "Elements", "ElementsByJS", "ElementsX", "Equal", "Eval", "Evaluate", "Focus", "Frame", "GetContext", "GetSessionID", "GetXPath", "HTML", "Has", "HasR", "HasX", "Hover", "Input", "InputColor", "InputTime", "Interactable", "KeyActions", "Matches", "MoveMouseOut", "MustAttribute", "MustBackgroundImage", "MustBlur", "MustCanvasToImage", "MustClick", "MustContainsElement", "MustDescribe", "MustDisabled", "MustDoubleClick", "MustElement", "MustElementByJS", "MustElementR", "MustElementX", "MustElements", "MustElementsByJS", "MustElementsX", "MustEqual", "MustEval", "MustFocus", "MustFrame", "MustGetXPath", "MustHTML", "MustHas", "MustHasR", "MustHasX", "MustHover", "MustInput", "MustInputColor", "MustInputTime", "MustInteractable", "MustKeyActions", "MustMatches", "MustMoveMouseOut", "MustNext", "MustParent", "MustParents", "MustPrevious", "MustProperty", "MustRelease", "MustRemove", "MustResource", "MustScreenshot", "MustScrollIntoView", "MustSelect", "MustSelectAllText", "MustSelectText", "MustSetFiles", "MustShadowRoot", "MustShape", "MustTap", "MustText", "MustType", "MustVisible", "MustWait", "MustWaitEnabled", "MustWaitInteractable", "MustWaitInvisible", "MustWaitLoad", "MustWaitStable", "MustWaitVisible", "MustWaitWritable", "Next", "Overlay", "Page", "Parent", "Parents", "Previous", "Property", "Release", "Remove", "Resource", "Screenshot", "ScrollIntoView", "Select", "SelectAllText", "SelectText", "SetFiles", "ShadowRoot", "Shape", "Sleeper", "String", "Tap", "Text", "Timeout", "Type", "Visible", "Wait", "WaitEnabled", "WaitInteractable", "WaitInvisible", "WaitLoad", "WaitStable", "WaitStableRAF", "WaitVisible", "WaitWritable", "WithCancel", "WithPanic"},
		},
	}
}

// --------------- proxy for github.com/go-rod/rod.CDPClient ---------------
type P__root_go_src_rod_CDPClient struct {
	Object interface{}
	Call_  func(_proxy_obj_ interface{}, ctx context.Context, sessionID string, method string, params interface{}) ([]byte, error)
	Event_ func(interface{}) <-chan *cdp.Event
}

func (P *P__root_go_src_rod_CDPClient) Call(ctx context.Context, sessionID string, method string, params interface{}) ([]byte, error) {
	return P.Call_(P.Object, ctx, sessionID, method, params)
}
func (P *P__root_go_src_rod_CDPClient) Event() <-chan *cdp.Event {
	return P.Event_(P.Object)
}
