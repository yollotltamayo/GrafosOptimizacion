interface ResponseFloydW {
  cambios: Array<Change>,
  iteraciones: Array<Array<Vertices>>,
  nodos: Array<string>,
}

interface ResponseFlujoMaximo {
  Flujo: number,
  Data: Array<Path>,
}

interface ResponseCPM {
  actividades: Array<Actividad>,
  rutaCritica: Array<string>,
  duracionTotal: number,
}

interface coor {
  row : number,
  col : number,
}

interface ResponseDijkstra {
  bases : Array<string>,
  destino: string,
  origen: string,
  peso : number,
  tabla: Array<Array<string>>,
  coords: Array<coor>,
}

interface ResponsePERT {
  rutaCritica: Array<string>,
  estimaciones: Array<number>,
  varianzas: Array<number>,
  sumaVariazas: number,
  media: number,
  cpm: ResponseCPM,
}

interface Actividad {
  nombre: string,
  duracion: number,
  sucesores: Array<string>,
  proximoL: number,
  proximoR: number,
  lejanoL: number,
  lejanoR: number,
}

interface Path {
  data: Array<Vertices>,
  camino: string,
}

// TODO  Vertices -> Vertice
interface Vertices {
  origen: string,
  destino: string,
  peso: number,
}

interface VerticePERT {
  origen: string,
  destino: string,
  optimista: number,
  probable: number,
  pesimista: number,
}

interface Change {
  iteracion: number,
  origen: string,
  destino: string,
}

interface VerticeCompresion{
  actividad:   string,
  predecesora: string,

  pesoNormal: number,
  costoNormal:number,

  pesoUrgente: number,
  costoUrgente: number,
}

interface CompresionData {
  tiempoObjetivo: number,
  actividades: Array<VerticeCompresion>,
}

interface ResponseCompresion {
  costoTiempo: Array<number>,
  iteraciones: Array<ResponseCPM>,
  actividadesComprimidas: Array<string>,
  costoActual: Array<number>,
}

interface ResponseKruskal {
  grafo: Array<Vertices>,
  arbol: Array<number>,
  peso: number,
  slides: string,
}
