"""Se define como puente local de un grafo a una arista que une dos vértices sin adyacentes en común. Implementar un
algoritmo que reciba un grafo y devuelva una lista con todos los puentes locales. Indicar y justificar la complejidad del
algoritmo implementado. Recomendación: no te la rebusques."""


def puentes_locales(grafo):
    lista = []
    for v in grafo.obtener_vertices():
        vecinos_v = grafo.adyacentes(v)
        for w in grafo.adyacentes(v):
            vecinos_w = grafo.adyacentes(w)
            for x in vecinos_w:
                if x in vecinos_v:
                    break
                lista.append((v, w))
    return lista


[10, 5, 8, 14, 4]

[10, 5, 8, 14, 7, 4]
