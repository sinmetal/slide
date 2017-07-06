ParDo.withOutputTags(packageObjects, TupleTagList.of(invalidRecord)).of(new DoFn<String, KV<String, String>>() {
    private static final long serialVersionUID = 1L;

	@Override
	public void processElement(ProcessContext c) {
	    Tokenizer tokenizer = new Tokenizer();
		List<Token> tokens = tokenizer.tokenize(c.element());
		for (Token token : tokens) {
		    if (token.getAllFeaturesArray().length < 1 || token.getAllFeaturesArray()[0].equals("記号")) {
                // invalid dataをside outputに逃がす
			    c.sideOutput(invalidRecord, token.getSurface());
			} else {
			    c.output(KV.of(token.getAllFeaturesArray()[0], token.getSurface()));
			}
		}
	}
}).named("Tokenizer"));