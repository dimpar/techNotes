### StringBuffer vs StringBuilder
- StringBuffer is thread safe or synchornized. Multiple thread can't access simultaneously.
StringBuilder is not thread safe.
- StringBugger methods are synchronized. StringBuilders are not.
- StringBuffer's performace is low. StringBuilder's performance is high.

### hashCode() and equals()
You must override hashCode() in every class that overrides equals(). Failure to do so will result in a violation of the general contract for Object.hashCode(), which will prevent your class from functioning properly in conjunction with all hash-based collections, including HashMap, HashSet, and Hashtable.