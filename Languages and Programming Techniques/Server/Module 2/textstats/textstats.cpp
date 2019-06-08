#include <iostream>
#include <set>
#include "textstats.hpp"

using namespace std;

void get_tokens(
        const string &s,
        const unordered_set<char> &delimiters,
        vector<string> &tokens
) {
    string current_token;
    for (auto ch : s) {
        if (delimiters.find(ch) == delimiters.end()) {
            current_token.push_back(tolower(ch));
        } else if (!current_token.empty()) {
            tokens.push_back(current_token);
            current_token = "";
        }
    }
    if (!current_token.empty()) {
        tokens.push_back(current_token);
        }
}

void get_type_freq(
        const vector<string> &tokens,
        map<string, int> &freqdi
) {
    for (auto token : tokens) {
        if (freqdi.find(token) != freqdi.end()) {
            ++freqdi[token];
        } else {
            freqdi.insert(make_pair(token, 1));
        }
    }
}

void get_types(
        const vector<string> &tokens,
        vector<string> &wtypes
) {
    set<string> uniq_words = set<string>();
    for (auto word : tokens) {
        uniq_words.insert(word);
    }
    for (auto uword : uniq_words) {
        wtypes.push_back(uword);
    }
}

void get_x_length_words(
        const vector<string> &wtypes,
        int x,
        vector<string> &words
) {
    for (auto word : wtypes) {
        if (word.size() >= x) {
            words.push_back(word);
        }
    }
}

void get_x_freq_words(
        const map<string, int> &freqdi,
        int x,
        vector<string> &words
) {
    for (auto pair : freqdi) {
        if (pair.second >= x) {
            words.push_back(pair.first);
        }
    }
}

void get_words_by_length_dict(
        const vector<string> &wtypes,
        map<int, vector<string> > &lengthdi
) {
    for (auto uword : wtypes) {
        lengthdi[uword.size()].push_back(uword);
    }
}