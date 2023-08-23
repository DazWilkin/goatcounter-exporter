# Test

```bash
export $(cat .env.test | xargs)
```

## `/me`

```bash
curl \
--get \
--header "Authorization: Bearer ${TOKEN}" \
https://pretired.goatcounter.com/api/v0/me \
| jq -r .
```

## `/stats`

```bash
curl \
--get \
--header "Authorization: Bearer ${TOKEN}" \
https://pretired.goatcounter.com/api/v0/stats/total \
| jq -r .

```