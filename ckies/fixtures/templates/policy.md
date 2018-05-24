{{ define "table" }}

| Name | Purpose | Term |
| ---- | ------- | ---- |
{{- range . }}
| {{ .Name }} | {{ .Info }} | {{ .Expires }} |
{{- end }}

{{- end -}}

# Cookie Policy

With this Cookie Policy we would like to inform you about the use of cookies and other similar storage technologies (hereafter called "Cookies") on our website.

## What are Cookies?

Cookies are small text files placed by your browser of your device in order to store certain information. When visiting our website again with the same device, the information stored in the Cookies will either be sent back to our website ("First Party Cookies") or to another website to which the Cookie belongs ("Third Party Cookies").

Through the stored and sent back information the respective website detects that you have previously visited it before with the browser of your device. We use this information to design and display the website optimally to you according to your preferences. During this process only the Cookie itself is identified on your device. A storage of personal data exceeding this only takes place after you have provided your explicit consent or if this is strictly necessary for enabling you the use of the service you requested.

## Consent to the use of Cookies

Cookies that are not strictly necessary to provide the services on our website will only be used after you have given your consent to their use. As soon as you use our website actively (by continuing to surf on our website after the Cookie Banner has been displayed to you), you give your consent to the use of Cookies.

You may customize your [Cookie Settings]({{ .Links.Settings }}) any time by activating or deactivating the different Cookie category. For more information on how to generally deactivate Cookies on your browser (including Strictly Necessary Cookies), see Section 4 below.

## Types of Cookies

Depending on its purpose and function we divide Cookies into the following categories:

 - Strictly Necessary Cookies;
 - Functional Cookies;
 - Performance Cookies; and
 - Marketing Cookies

### Strictly Necessary Cookies

Strictly Necessary Cookies secure certain functions, without which you would not be able to use our website as intended. These Cookies are First Party Cookie, as they are only used and set by us. This means that all information, which is stored in the Cookies, is sent back to our website.

Strictly Necessary Cookies are used for instance to keep you logged-in as a user when accessing different sub-sites of our website, so that you do not have to log-in every time you access a new site.

{{ if (and (not .Options.ForceTables) (not .Cookies.Necessary)) -}}
  Our website does not use Strictly Necessary Cookies.
{{- else -}}
  Our website uses the following Strictly Necessary Cookies:

  {{- template "table" .Cookies.Necessary -}}
{{- end }}

The use of Strictly Necessary Cookies on our website is possible without your consent. Therefore, Strictly Necessary Cookies cannot be separately activated or deactivated. This function is only available for Cookies that require your prior consent, thus Functional, Performance and Marketing Cookies. However, you will always have the possibility to generally deactivate all Cookies in your browser settings.

### Functional Cookies

Functional Cookies enable our website to store information previously provided by you (such as registered name or selected language) and to offer you improved and personalized features based thereon. These Cookies only collect and store anonymized information, so that they cannot track your movements on other websites.

{{ if (and (not .Options.ForceTables) (not .Cookies.Functional)) -}}
  Our website does not use Functional Cookies.
{{- else -}}
  Our website uses the following Functional Cookies:

  {{- template "table" .Cookies.Functional -}}
{{- end }}

You may object to the use of Functional Cookies at any time by customizing your [Cookie Settings]({{ .Links.Settings }}) accordingly.

### Performance Cookies

Performance Cookies collect information on how our website is used in order for us to improve its appeal, content and functionality. These Cookies help us determine, for example, which sub-sites of our website are accessed and what content the users are most interested in. In particular, we register the number of access to each site and sub-sites, the time spent on our website, the sequence of the sites visited, which search terms led you to us, the country, region and city from which the website is accessed as well as the proportion of mobile devices accessing our website. Moreover, we register all movements, clicks and scrolling with the computer mouse in order to determine which sections of our website the users are most interested in. As a consequence, we are able to adjust the content of our website to meet the specific needs of our users and optimize our service. IP addresses, which have to be transmitted due to technical reasons, are automatically anonymized and therefore no link to individual users may be drawn by us.

We reserve the right to use the information gained through the anonymized analysis of the user behavior via Cookies also in order to display to you specific advertisement for certain products of ours on our own websites. We believe that you as a user profit from this because we display advertisements and content we assume meet your interests on the basis of your surfing behavior. On the contrary, less random advertisement or less interesting content for you will be displayed.

{{ if (and (not .Options.ForceTables) (not .Cookies.Performance)) -}}
  Our website does not use Performance Cookies.
{{- else -}}
  Our website uses the following Performance Cookies:

  {{- template "table" .Cookies.Performance -}}
{{- end }}

You may object to the use of Performance Cookies at any time by customizing your [Cookie Settings]({{ .Links.Settings }}) accordingly.

### Marketing Cookies / Third Party Cookies

Marketing Cookies are set by external advertising companies (Third Party Cookies) and will be used to gather information about the user’s visited websites in order to create targeted advertising to the user.

{{ if (and (not .Options.ForceTables) (not .Cookies.Marketing)) -}}
  Our website does not use Marketing Cookies.
{{- else -}}
  Our website uses the following Marketing Cookies:

  {{- template "table" .Cookies.Marketing -}}
{{- end }}

You may object to the use of Marketing Cookies at any time by customizing your [Cookie Settings]({{ .Links.Settings }}) accordingly.

## Administration and deletion of cookies

Furthermore, you may change the settings of your internet browser to prevent the storage of Cookies on your device in general or that you will be asked every time, whether you agree to the setting and use of Cookies. Cookies that have been set once can also be deleted at any time. How all this works in detail can be found in your browser’s help function.

Please note that a general deactivation of the use of Cookies might possibly lead to a restriction of our website´s functionality.

Cookie Policy as of May 2018
