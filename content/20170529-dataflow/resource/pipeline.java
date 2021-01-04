public static void main(String[] args) {
		DatastoreToBigQueryOptions options = PipelineOptionsFactory.fromArgs(args).withValidation()
				.as(DatastoreToBigQueryOptions.class);
		Pipeline p = Pipeline.create(options);
		...
        // Input from Datastore
		PCollection<Entity> entities = p
				.apply(DatastoreIO.v1().read().withProjectId(options.getProject()).withQuery(query));
		PCollection<TableRow> rows = entities.apply(new EntityToTableRow());
        ...
        // Output to BigQury
		rows.apply(BigQueryIO.Write.named("Write").to("training-topgate-dev:output.output_table").withSchema(schema)
				.withWriteDisposition(BigQueryIO.Write.WriteDisposition.WRITE_TRUNCATE)
				.withCreateDisposition(BigQueryIO.Write.CreateDisposition.CREATE_IF_NEEDED));

		p.run();
	}
