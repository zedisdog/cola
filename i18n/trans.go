package i18n

var defaultTranslator = New("zh_CN")

//Translator struct holds translate data
type Translator struct {
	trans       map[string]map[string]string
	defaultLang string
}

func New(defaultLang string) *Translator {
	return &Translator{
		trans:       make(map[string]map[string]string),
		defaultLang: defaultLang,
	}
}

func Set(name string, trans string) {
	defaultTranslator.Set(name, trans)
}
func (t *Translator) Set(name string, trans string) {
	t.SetByLang(name, t.defaultLang, trans)
}

func SetByLang(name, lang, trans string) {
	defaultTranslator.SetByLang(name, lang, trans)
}
func (t *Translator) SetByLang(name string, lang string, trans string) {
	if _, ok := t.trans[name]; !ok {
		t.trans[name] = map[string]string{
			lang: trans,
		}
	} else {
		t.trans[name][lang] = trans
	}
}

func TransError(err error) string {
	return defaultTranslator.TransError(err)
}
func (t *Translator) TransError(err error) string {
	return t.TransErrorByLang(err, t.defaultLang)
}

func TransErrorByLang(err error, lang string) string {
	return defaultTranslator.TransErrorByLang(err, lang)
}
func (t *Translator) TransErrorByLang(err error, lang string) string {
	return t.TransByLang(err.Error(), lang)
}

func Trans(name string) string {
	return defaultTranslator.Trans(name)
}
func (t *Translator) Trans(name string) string {
	return t.TransByLang(name, t.defaultLang)
}

func TransByLang(name, lang string) string {
	return defaultTranslator.TransByLang(name, lang)
}
func (t *Translator) TransByLang(name string, lang string) string {
	if _, ok := t.trans[name]; !ok {
		return ""
	} else {
		return t.trans[name][lang]
	}
}
