/*
	LENGUAJES Y AUTOMATAS I
	HERNANDEZ CONTRERAS UZZIEL
	13211489
*/

package main

func email() string {
	return `[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}`
}

func url() string {
	return `(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[(com)|(org)|(edu)|(net)|(biz)]{1}`
}

func tipoDato() string {
	return `(negrita|blanca|redonda|corchea|fusa|semifusa){1}`
}

func repeticion() string {
	return `coda`
}

func clase() string {
	return `Pentagrama`
}

func pausa() string {
	return `silencio`
}

func modificador() string {
	return `(bemol|sostenido|sustraido|becuadro){1}`
}

func asigna() string {
	return `=`
}

func elMain() string {
	return `partitura`
}

func selector() string {
	return `armadura`
}

func objeto() string {
	return `compas`
}

func reservadas() string {
	return `(pp|disminuido|aumentado|sustraido|septima|menor|
	|quinta|ff|pf|ppp|fff|lento|crescendo|disminuendo|fine|tempo|larghissimo|largo|
	|moderatto|grave|largheto|andantino|adagio|tranquilo|affetuosso|andante|espressivo|
	|allegrato|allegro|vivace|allegrissimo|prestissimo|presto|vivacissimo|grazioso|
	|tranquillamente|stretto|stringendo|accelerando|affrettando|rallentando|ritardando|
	|ritenuto|Apiacere|Acapriccio|Adlibitum|rubato|Atempo|sostenuto|morendo|Nontroppo|
	|Con moto|molto|Tempodi|quasi|assai|Tempogiusto|
	|Tempoprimo|Abeneplacito|Acapella|acarezzevole|acento|acceso|acciaccatura|accompagnato|Acorde roto|
	|adagietto|adagissimo|A due|affanato|agile|agitato|altissimo|amabile|amoroso|anacrusa|antifonia|apaise|
	|tiempo|ritmo|Mezzo piano|Mezzoforte|fortissimo|bending|muteo|tapping|puntillo|tresillo|caesura|
	|ligadura|legato|glissando|picado|staccato|marcato|pizzicato|sincopa|A capella|trino|mordente|
	|grupeto|apoyatura|ottava alta|quindicesima|alta|ritornelo|fermata|adagio|subdominante|dominante|
	|tablatura|armonia|melodia|escala|percusion){1}`
}

func apuntadores() string {
	return `(arpegio|Bassocontinuo|bellicoso|cadenza|calderon|canto|canon|
	|cantabile|compasillo|Doblescuerdas|drammatico|downtempo|falsetto|festivamente|forte|freddo|
	|Gp|octava|pedale|semitono|silenzio|solenne|soprano|tenor|ternario|tremolo|vibrato|voce|virtuoso|
	|Unacorda|Trecorda|tranquillo|tenuto|tacet|subito){1}`
}

func id() string {
	return `[a-z]+[0-9]*`
}

func opRel() string {
	return `(([<>]){1}(=)?|(==)|(!=))`
}

func agDer() string {
	return `(\{)`
}

func agIzq() string {
	return `(\})`
}

func numeros() string {
	return `((-)?[0-9]+)`
}

func parDer() string {
	return `(\()`
}

func parIzq() string {
	return `(\))`
}

func opMat() string {
	return `(\+|\-|\*|/){1}`
}

func pregunta() string {
	return "clave"
}

func saltoLinea() string {
	return `^(\n)*$`
}

func bloque() string {
	return `(.)*`
}

func comentarios() string {
	return "(//)+(.)*$"
}
