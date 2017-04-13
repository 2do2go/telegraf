# Check DNS Domain Input Plugin

This input plugin will return how much time (in seconds) left for a domain paid period to expire.

### Configuration:

```
# Domain configuration given a list of domains
[[inputs.check_dns_domain]]
  #Domains
  domains = ["github.com"]
```

### Measurements & Fields:

- domains
    - time_to_expire (int) # seconds left for the domain expiration date

### Tags:

- All measurements have the following tags:
    - domain

### Example Output:

```
$ ./telegraf -config telegraf.conf -input-filter check_dns_domain -test
> domains,domain=google.com time_to_expire=6185474.476944118 1468864305580596685
```
