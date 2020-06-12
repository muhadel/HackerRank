# Complete the matchingStrings function below.
def matchingStrings(strings, queries):
    frequency = {}
    results = []
    for str in strings:
        if str in frequency:
            frequency[str] += 1
        else:
            frequency[str] = 1

    for query in queries:
        results.append(query in frequency and frequency[query] or 0)
     
    return results

if __name__ == '__main__':
    strings = ["ab", "ab", "abc"]
    queries = ["ab", "abc", "bc"]
    print(matchingStrings(strings, queries))