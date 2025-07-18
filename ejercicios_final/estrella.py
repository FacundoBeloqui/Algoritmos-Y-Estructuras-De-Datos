"""Implementar un algoritmo que reciba un grafo no dirigido y determine si el mismo tiene forma de estrella. Es decir,
si todos los vértices, salvo 1, se conectan al mismo vértice, mientras ese único vértice se conecta con todos los demás.
Indicar y justificar la complejidad del algoritmo si se implementara el grafo con una lista de adyacencia (diccionario de
diccionarios), y también si se hiciera con una matriz de adyacencia."""


def es_estrella(grafo):
    for v in grafo:
        if len(grafo.adyacentes(v)) == len(grafo) - 1:
            contador = 0
            for w in grafo.adyacentes(v):
                if len(grafo.adyacentes(w)) == 1 and v in grafo.adyacentes(w):
                    contador += 1
            if contador == len(grafo) - 1:
                return True
    return False
