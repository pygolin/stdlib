
PYGOLINC = grumpc

all: marshal/module.go \
pickle/module.go \
sys/module.go

marshal/module.go: src/marshal.py
	$(PYGOLINC) -m marshal $^ > $@
	gofmt -w $@

pickle/module.go: src/pickle.py
	$(PYGOLINC) -m pickle $^ > $@
	gofmt -w $@

sys/module.go: src/sys.py
	$(PYGOLINC) -m sys $^ > $@
	gofmt -w $@

