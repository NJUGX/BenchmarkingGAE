package fibo;

import java.io.IOException;
import javax.servlet.ServletException;

import com.google.appengine.api.datastore.DatastoreService;
import com.google.appengine.api.datastore.DatastoreServiceFactory;
import com.google.appengine.api.datastore.Entity;
import com.google.appengine.tools.development.testing.LocalDatastoreServiceTestConfig;
import com.google.appengine.tools.development.testing.LocalServiceTestHelper;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import java.util.Date;
import org.junit.After;
import org.junit.Before;


@RunWith(JUnit4.class)
public class HelloAppEngineTest {

	private final LocalServiceTestHelper helper = new LocalServiceTestHelper(
			new LocalDatastoreServiceTestConfig().setDefaultHighRepJobPolicyUnappliedJobPercentage(0));

	private DatastoreService datastore;

	@Before
	public void setUp() {
		helper.setUp();
		datastore = DatastoreServiceFactory.getDatastoreService();
	}

	@After
	public void tearDown() {
		helper.tearDown();
	}

	@Test
	public void test() throws IOException, ServletException {

		MockHttpServletResponse response = new MockHttpServletResponse();
//		MockHttpServletRequest request = new MockHttpServletRequest();
		new HelloAppEngine().doGet(null, response);
		Assert.assertEquals("text/plain", response.getContentType());
		Assert.assertEquals("UTF-8", response.getCharacterEncoding());
	}
}
