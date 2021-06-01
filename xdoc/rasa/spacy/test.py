import spacy

parser = spacy.load('zh_core_web_trf')

doc = parser('打印错误测试。')
for token in doc[:17]:
    print("{0}\t{1}\t{2}\t{3}\t{4}\t{5}\t{6}\t{7}\t{8}\t{9}\t{10}\t{11}\n".format(
        token.text,   # 文本
        token.idx,  # 索引值（即在原文中的定位）
        token.lemma_,  # 词元(lemma)
        token.head,   # 当前Token的Parent Token，从语法关系上来看，每一个Token都只有一个Head。
        token.dep_, # 依存关系
        token.children, # 语法上的直接子节点
        token.ancestors, # 语法上的父节点
        token.is_punct, # 是否为标点符号
        token.is_space,  # 是否为空格
        token.shape_,  # 字个数用x表示，如：两个字就是xx
        token.pos_,  # 词性
        token.tag_  # 标记
    ))
