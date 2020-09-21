/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package helper

import (
	"bytes"
	"errors"
	"github.com/apimatic/unirest-go"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

/**
 * Replaces template parameters in the given url
 * @param	queryBuilder    The query string builder to replace the template parameters
 * @param	parameters		The parameters to replace in the url
 * @return	Returns url replaced with template values
 */
func ToString(data interface{}, dvl string) string {
	return unirest.ToString(data, dvl)
}
func AppendUrlWithTemplateParameters(queryBuilder string, parameters map[string]interface{}) (string, error) {
	//perform parameter validation
	if len(queryBuilder) < 1 {
		return queryBuilder, errors.New("Given value for parameter \"queryBuilder\" is invalid.")
	}
	if len(parameters) < 1 {
		return queryBuilder, nil
	}

	//iterate and append parameters
	for key, value := range parameters {
		//ignore null values
		if value == nil {
			continue
		}

		replaceValue := ""

		//load parameter value
		if str, ok := value.(string); ok {
			replaceValue = str
		} else if strA, ok := value.([]string); ok {
			replaceValue = strings.Join(strA, "/")
		} else if strA, ok := value.([]int64); ok {
			for i := 0; i < len(strA); i++ {
				if i == 0 {
					replaceValue += unirest.ToString(strA[i], "")
				} else {
					replaceValue += "/" + unirest.ToString(strA[i], "")
				}
			}

		} else {
			replaceValue = url.QueryEscape(ToString(value, ""))
		}

		//to replace
		toReplace := "{" + key + "}"

		//find the template parameter and replace it with its value
		queryBuilder = strings.Replace(queryBuilder, toReplace, replaceValue, -1)
	}
	return queryBuilder, nil
}

/**
 * Appends the given set of parameters to the given query string
 * @param	string	queryBuilder	The query url string to append the parameters
 * @param	map     parameters	The parameters to append
 * @return	void
 */
func AppendUrlWithQueryParameters(queryBuilder string, parameters map[string]interface{}) (string, error) {
	//perform parameter validation
	if len(queryBuilder) < 1 {
		return queryBuilder, errors.New("Given value for parameter \"queryBuilder\" is invalid.")
	}
	if len(parameters) < 1 {
		return queryBuilder, nil
	}
	//does the query string already has parameters
	hasParams := (strings.Index(queryBuilder, "?") > 0)

	var buf bytes.Buffer
	buf.WriteString(queryBuilder)

	//iterate and append parameters
	for key, value := range parameters {
		//ignore null values
		if value == nil {
			continue
		}

		var toAppend string

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			toAppend = arrayToQuery(key, value)
		} else {
			strVal := unirest.ToString(value, "")
			if len(strVal) > 0 {
				toAppend = url.QueryEscape(key) + "=" + url.QueryEscape(strVal)
			}
		}

		//have something to append?
		if len(toAppend) < 1 {
			continue
		}

		//if already has parameters, use the &amp; to append new parameters
		if hasParams {
			buf.WriteString("&")
		} else {
			buf.WriteString("?")
		}

		//indicate that now the query has some params
		hasParams = true

		//append value
		buf.WriteString(toAppend)
	}

	return buf.String(), nil
}

/**
 * Converts an array type key values pair into a query string
 * @param	key		The key for the pair
 * @param	values	The array of values for the query
 */
func arrayToQuery(key string, data interface{}) string {
	var buf bytes.Buffer
	values := reflect.ValueOf(data)

	for index := 0; index < values.Len(); index++ {
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		value := unirest.ToString(values.Index(index).Interface(), "")
		buf.WriteString(url.QueryEscape(value))
		buf.WriteByte('&')
	}
	s := buf.String()
	if len(s) < 1 {
		return ""
	}
	return s[0 : len(s)-1]
}

/**
 * Validates and processes the given Url
 * @param    url     The given Url to process
 * @return   Pre-processed Url as string
 */
func CleanUrl(url string) (string, error) {
	//perform parameter validation
	if len(url) < 1 {
		return "", errors.New("Invalid Url.")
	}
	//ensure that the urls are absolute
	isMatch, err := regexp.MatchString("^(https?://[^/]+)", url)
	if err != nil || !isMatch {
		return "", errors.New("Invalid Url format.")
	}
	var query, parameters string
	//get the http protocol match
	protocol := url[0:7]
	if strings.Contains(url, "?") {
		query = url[7:strings.Index(url, "?")]
		parameters = url[strings.Index(url, "?"):len(url)]
	} else {
		query = url[7:len(url)]
		parameters = ""
	}

	//remove redundant forward slashes
	re := regexp.MustCompile("//+")
	query = re.ReplaceAllString(query, "/")
	url = protocol + query + parameters

	//return process url
	return url, nil
}
