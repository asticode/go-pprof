# PPROF

This package allows a smooth use of the pprof package

# Usage

- Add the following in your code :

        c, err := pprof.Profile()
        if err != nil {
            // Process error
        }
        defer c.Close()

- Next time you call your binary, either add the flag `-profile-cpu`, `-profile-mem` or both

Results will be written in `profile.cpu` and `profile.mem` files.
