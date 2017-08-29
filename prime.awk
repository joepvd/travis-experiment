#!/usr/bin/env gawk

BEGIN {
    n=ARGV[1]; delete ARGV[1]
    proc=ARGV[2]; delete ARGV[2]
    for (i=2; i<=n; i++)
        p[i]=0
    print proc, "Done initializing"
    for (i=2; i<=n; i++) {
        if (i % 10 == 0)
            print proc, "processing", i
        if (p[i]==0) {
            k=i*i
            for (c=0; k<=(n*n); k = i*i + ++c * i) {
                p[k]=1
            }
        }
    }
    for (prime=2; prime in p; prime++)
        if (p[prime]==0) print proc, prime
}


