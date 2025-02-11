# flotte

**Erläuterung der Zeichnung:**

* **Start: Input Video (A):**  Der Prozess beginnt mit dem Eingabevideo, das transkodiert werden soll.
* **Video Segmentierung (B):** Das Eingabevideo wird in mehrere Segmente aufgeteilt. Dies ermöglicht die parallele Verarbeitung durch die Map-Nodes.
* **Segment 1, Segment 2, Segment N (C, D, E):** Repräsentieren die einzelnen Video-Segmente, die erzeugt wurden.
* **Map Node 1, Map Node 2, Map Node N: Map Task (F1, F2, FN):**  Stellen die verschiedenen Knoten im MapReduce-Cluster dar, die die Map-Aufgaben ausführen. Jeder Knoten erhält ein oder mehrere Video-Segmente zur Verarbeitung.
* **Dekodierung (G1, G2, GN):**  Innerhalb jedes Map-Nodes wird das zugewiesene Video-Segment zunächst dekodiert.
* **Transkodierung (H1, H2, HN):**  Nach der Dekodierung erfolgt die eigentliche Transkodierung des Segments in das gewünschte Zielformat.
* **Transkodiertes Segment 1, Transkodiertes Segment 2, Transkodiertes Segment N (I1, I2, IN):**  Die Ergebnisse der Map-Phase. Jeder Map-Node gibt ein transkodiertes Video-Segment aus.
* **Reduce Node: Reduce Task (J):**  Ein zentraler Knoten, der die Reduce-Aufgabe übernimmt. Er sammelt die transkodierten Segmente von den Map-Nodes.
* **Segment Aggregation (K):**  Der Reduce-Node führt die transkodierten Segmente in der korrekten Reihenfolge zusammen.
* **Finale Video Enkodierung (L):**  Nach der Aggregation kann eine finale Enkodierung des gesamten Videos erfolgen, falls notwendig.
* **Ende: Transkodiertes Video (M):**  Der Prozess endet mit der Ausgabe des vollständig transkodierten Videos.

**Mermaid Code:**

```mermaid
graph LR
    A[Start: Input Video] --> B{Video Segmentierung};
    B --> C(Segment 1);
    B --> D(Segment 2);
    B --> E(Segment N);
    C --> F1[Map Node 1: Map Task];
    D --> F2[Map Node 2: Map Task];
    E --> FN[Map Node N: Map Task];
    F1 --> G1{Dekodierung};
    F2 --> G2{Dekodierung};
    FN --> GN{Dekodierung};
    G1 --> H1{Transkodierung};
    G2 --> H2{Transkodierung};
    GN --> HN{Transkodierung};
    H1 --> I1[Transkodiertes Segment 1];
    H2 --> I2[Transkodiertes Segment 2];
    HN --> IN[Transkodiertes Segment N];
    I1 & I2 & IN --> J[Reduce Node: Reduce Task];
    J --> K{Segment Aggregation};
    K --> L{Finale Video Enkodierung};
    L --> M[Ende: Transkodiertes Video];

    style F1,F2,FN fill:#f9f,stroke:#333,stroke-width:2px
    style J fill:#ccf,stroke:#333,stroke-width:2px
```
