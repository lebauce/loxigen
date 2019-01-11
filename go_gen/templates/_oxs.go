:: # Copyright 2013, Big Switch Networks, Inc.
:: # Copyright 2018, Red Hat, Inc.
:: #
:: # LoxiGen is licensed under the Eclipse Public License, version 1.0 (EPL), with
:: # the following special exception:
:: #
:: # LOXI Exception
:: #
:: # As a special exception to the terms of the EPL, you may distribute libraries
:: # generated by LoxiGen (LoxiGen Libraries) under the terms of your choice, provided
:: # that copyright and licensing notices generated by LoxiGen are not altered or removed
:: # from the LoxiGen Libraries and the notice provided below is (i) included in
:: # the LoxiGen Libraries, if distributed in source code form and (ii) included in any
:: # documentation for the LoxiGen Libraries, if distributed in binary form.
:: #
:: # Notice: "Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler."
:: #
:: # You may not use this file except in compliance with the EPL or LOXI Exception. You may obtain
:: # a copy of the EPL at:
:: #
:: # http://www.eclipse.org/legal/epl-v10.html
:: #
:: # Unless required by applicable law or agreed to in writing, software
:: # distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
:: # WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
:: # EPL for the specific language governing permissions and limitations
:: # under the EPL.

func (self *${ofclass.goname}) GetOXSName() string {
	return "${ofclass.name[7:]}"
}

func (self *${ofclass.goname}) GetOXSValue() interface{} {
	return self.Value
}

func (self *${ofclass.goname}) MarshalJSON() ([]byte, error) {
	var value interface{} = self.GetOXSValue()
	switch t := value.(type) {
	case net.HardwareAddr:
		value = t.String()
	case net.IP:
		value = t.String()
	default:
		if s, ok := t.(fmt.Stringer); ok {
			value = s.String()
		} else {
			value = t
		}
	}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("{\"Type\":\"%s\",\"Value\":%s}", self.GetOXSName(), string(jsonValue))), nil
}