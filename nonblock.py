import os
import sys
import fcntl

flags_stdout = fcntl.fcntl(sys.stdout, fcntl.F_GETFL)
fcntl.fcntl(sys.stdout, fcntl.F_SETFL, flags_stdout&~os.O_NONBLOCK)

flags_stderr = fcntl.fcntl(sys.stderr, fcntl.F_GETFL)
fcntl.fcntl(sys.stderr, fcntl.F_SETFL, flags_stderr&~os.O_NONBLOCK)

