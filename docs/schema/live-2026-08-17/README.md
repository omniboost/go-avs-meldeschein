# Live schema snapshot — 2026-08-17

Fetched directly from AVS's production endpoint (`https://meldeschein.avs.de/meldeschein-ws-prod/JMeldescheinWebservices`)
via the standard JAX-WS `?wsdl` / `?xsd=N` introspection query params, authenticated
with the `guestline_omniboost` Basic Auth credentials shared by the guestline-avs
pipeline configs. Read-only metadata request, no booking/cancel operation invoked.

This supersedes the bundled `../meldeschein_*.xsd` files one directory up, which are
version 6.3 (2018) and predate `email`, `digit_gastkart`, `jkk`, `manueller_betrag`,
`authentifizierung` and `weitere_kosten` on `cMs`. It also corrects an assumption from
the v7.6 prose spec PDF: that document's booking-response *examples* show
`email`/`digit_gastkart` placed right after `geburtsdatum`, but the actual schema here
places them after `abrechnungstatusid`, `zusatzfelder` and `indiv-zusatzfelder` instead
— see IN-4787.

`schemaLocation`/`xsd:import` URLs have been rewritten from the live absolute
`https://meldeschein.avs.de:443/...` targets to relative paths within this directory
so the four files resolve against each other standalone.
