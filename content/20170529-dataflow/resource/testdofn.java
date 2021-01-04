@Test
public void testEntityToTableRowFn() {
	final String KIND = "SampleKind";
	final String ID = "hogeId";
	final String COLUMN_CONTENT = "content";
	final String COLUMN_CONTENT_VALUE = "hogeContent";

	DoFnTester<Entity, TableRow> extractEntityToTableRowFn = DoFnTester.of(new EntityToTableRowFn());

	...

	List<TableRow> rows = extractEntityToTableRowFn.processBatch(entity);
	Assert.assertThat(rows.size(), CoreMatchers.is(1));
	Assert.assertTrue(rows.get(0).containsKey(COLUMN_CONTENT));
	Assert.assertThat(rows.get(0).get(COLUMN_CONTENT).toString(), CoreMatchers.is(COLUMN_CONTENT_VALUE));
}