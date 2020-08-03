from math import log


def tf_idf(term, docs):
    """
    Term Frequency - Inverse Document Frequency Equation:
    tf-idf(t, d) = tf(t, d) * idf(t, d)
    """
    # Convert term to lower case
    term = term.lower()

    tfs = [term_frequency(term, doc) for doc in docs]
    idf = inverse_document_frequency(term, docs)

    tf_idfs = [tf.get(term) * idf for tf in tfs]

    return tf_idfs


def term_frequency(term, doc):

    """
    Term Frequency Equation:

    tf(t, d) = N(t, d), wherein tf(t, d) = term frequency for a term t in document d.
    N(t, d)  = number of times a term t occurs in document d
    """

    # Initialize a dictionary of normalized term frequencies
    normalized_tfs = {}

    # Split the document into a list of terms
    # Convert all keywords to lowercase
    normalize_term_frequency = doc.lower().split()

    # Count the # of times the term occurs in the document
    term_in_document = normalize_term_frequency.count(term)

    # Total number of terms in the document
    len_of_document = float(len(normalize_term_frequency))

    # Normalized Term Frequency
    normalized_tf = term_in_document / len_of_document
    normalized_tfs[term] = normalized_tf

    return normalized_tfs


def inverse_document_frequency(term, docs):
    num_docs_with_given_term = 0

    """
    Inverse Document Frequency Equation:

    idf(t) = log(N/ df(t))
    """
    # Iterate through all the documents
    for doc in docs:

        """ 
        Putting a check if a term appears in a document. 
        If term is present in the document, then  
        increment "num_docs_with_given_term" variable 
        """
        if term in doc.split():
            num_docs_with_given_term += 1

    if num_docs_with_given_term > 0:
        # Total number of documents
        total_num_docs = len(docs)

        # Calculating the IDF
        idf_val = log(float(total_num_docs) / num_docs_with_given_term)
        return idf_val
    else:
        return 0


doc_1 = "the cat in the hat"
doc_2 = "the dog in the fog"

keywords = "I am a man who likes dog and cat"

tfidfs = []
for word in keywords.lower().split():
    print(word)
    tfidfs.append(tf_idf(word, [doc_1, doc_2]))

print(tfidfs)
