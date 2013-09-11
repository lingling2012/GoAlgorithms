第六章 堆排序
===========

堆（Heap）
---

一个最大/最小堆支持以下操作：

MAX/MIN-HEAP:维护最大/最小堆，O(lgN)

BUILD-MAX/MIN-HEAP:建堆，O(N)

HEAP-SORT:堆排序，O(NlgN)



优先队列（Priority Queue）
-------

一个最大/最小优先队列支持以下操作：

INSERT(S, x): 把元素x插入集合S

MAXIMUM/MINIMUM(S): 返回S中具有最大/最小关键字的元素

EXTRACT-MAX/MIN(S): 去掉并返回S中具有最大/最小关键字的元素

INCREASE/DECREASE-KEY(S, x, k): 将元素x的关键字增加到/减少到k



最小优先队列可用于第23章的最小生成树Prim算法和第24章的最短路径Dijkstra算法。



利用堆可以实现一个优先队列:

MAX/MIN-HEAP-INSERT:

HEAP-MAXIMUM/MINIMUM:

HEAP-EXTRACT-MAX/MIN:

HEAP-INCREASE/DECREASE-KEY:

在一个包含N个元素的堆中，所有优先队列的操作都可以在O(lgN)时间内完成。



