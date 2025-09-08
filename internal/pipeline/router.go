//! M2 Log Analyzer AI
//! --------------------------------
//! License : Dual License
//!           - Apache 2.0 for open-source / personal use
//!           - Commercial license required for closed-source use
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Routes log events through pipeline rules defined in config.

package pipeline

type Router struct {
	rules []Rule
}

type Rule struct {
	Match   string
	Handler func(LogEvent)
}

func NewRouter(rules []Rule) *Router {
	return &Router{rules: rules}
}

func (r *Router) Route(event LogEvent) {
	for _, rule := range r.rules {
		// TODO: implement proper pattern matching
		if event.Level == rule.Match {
			rule.Handler(event)
		}
	}
}
