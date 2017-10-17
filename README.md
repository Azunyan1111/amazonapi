# amazonapi


###******* WARNING ******
Rewrite the following code.
You need to comment out.


github.com/svvu/gomws/gmws/base.go#NewMwsBase()
```github.com/svvu/gomws/gmws/base.go
	if config.SellerId == "" {
		return nil, fmt.Errorf("No seller id provided")
	}
	// ********** This code. *********
	//if config.AuthToken == "" {
	//	return nil, fmt.Errorf("No auth token provided")
	//}

	region := config.Region
	if region == "" {
		region = "US"
	}

	marketPlace, mError := marketplace.New(region)
	if mError != nil {
		return nil, mError
	}

	base := MwsBase{
		SellerId:      config.SellerId,
		AuthToken:     config.AuthToken,
		Region:        region,
		MarketPlaceId: marketPlace.Id,
		Host:          marketPlace.EndPoint,
		Version:       version,
		Name:          name,
		accessKey:     config.AccessKey,
		secretKey:     config.SecretKey,
	}
	return &base, nil
```
github.com/svvu/gomws/gmws/base.go#paramsToAugment()
```github.com/svvu/gomws/gmws/base.go
	clientInfo := map[string]string{
		"SellerId":         base.SellerId,
		// ********** This code. *********
		//"MWSAuthToken":     base.AuthToken,
		"SignatureMethod":  base.SignatureMethod(),
		"SignatureVersion": base.SignatureVersion(),
		"AWSAccessKeyId":   base.getCredential().AccessKey,
		"Version":          base.Version,
	}
	return clientInfo
```
