package timezone

func Code(c string) []Timezone {
	var z []Timezone
	for _, zone := range Locations {
		if zone.Code == c {
			z = append(z, zone)
		}
	}
	return z
}
