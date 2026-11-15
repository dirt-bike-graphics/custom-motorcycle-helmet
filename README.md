# Custom Motorcycle Helmet

Open reference guide and JSON catalog for **custom motorcycle helmet** shell styles, artwork zones and surface finishes. Covers all 33 helmet silhouettes used in the 3D AI designer at [Custom Motorcycle Helmet](https://custommotorcyclehelmet.com).

## Files in This Repository

| File | Contents |
|---|---|
| [index.html](index.html) | Interactive static reference table of 33 custom motorcycle helmet shells, artwork zones and finish presets |
| [styles.json](styles.json) | Machine-readable JSON catalog of all 33 helmet IDs, rider categories and shell profiles |

## Helmet Categories Covered (33 Styles)

* **Full face (12 styles):** Retro full face, Modern full face, Street full face, Race with tail wing, Sport with tail wing, Vintage full face, Retro with goggles, Rally full face, Slim full face, Car racing style, Race with chin spoiler and Rounded full face.
* **Motocross and dirt (7 styles):** Retro motocross, Retro scrambler without peak, Enduro shell, Downhill full face, Sharp motocross, Motocross and Motocross with goggles.
* **Adventure (2 styles):** Adventure dual sport and Street dual sport.
* **Open face (6 styles):** Retro open face, Jet with long shield, Retro open with peak, Open face, Classic open face and Open face with peak.
* **Half helmet (6 styles):** Half helmet, Half with peak, Shorty half, Beanie, Vintage half and German half.

## Using the JSON Data

```js
const res = await fetch("https://dirt-bike-graphics.github.io/custom-motorcycle-helmet/styles.json");
const { updated, total, styles } = await res.json();
const motocrossShells = styles.filter((s) => s.category === "Motocross");
```

## Official Designer and Guides

* Official site: [Custom Motorcycle Helmet](https://custommotorcyclehelmet.com)
* Live 3D editor: [Custom Motorcycle Helmet Designer](https://custommotorcyclehelmet.com/design)
* By helmet type: [Custom Full Face Motorcycle Helmets](https://custommotorcyclehelmet.com/custom-full-face-motorcycle-helmets) | [Custom Dirt Bike Helmets](https://custommotorcyclehelmet.com/custom-dirt-bike-helmets) | [Custom Half Helmets](https://custommotorcyclehelmet.com/custom-half-helmets)
* Finishes and decals: [Motorcycle Helmet Wraps](https://custommotorcyclehelmet.com/motorcycle-helmet-wraps) | [Motorcycle Helmet Decals](https://custommotorcyclehelmet.com/motorcycle-helmet-decals) | [Motorcycle Helmet Designs](https://custommotorcyclehelmet.com/motorcycle-helmet-designs)
