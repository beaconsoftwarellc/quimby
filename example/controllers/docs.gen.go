package controllers

// THIS IS A GENERATED FILE. DO NOT MODIFY
// api_docs.tmpl

const htmlDocs = `<!DOCTYPE html>
<html lang="en">
  <head>
	<!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  </head>
  <body>
    <div class="container">
      <h2>Endpoints</h2>
		
	  
	  
	  <div class="panel panel-success">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Get</b></div>
			  <div class="col-sm-10">/api/widgets</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetsGetInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetsGetInfo">
			  <p>returns a collection Widgets</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				
				
				
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">200</td>
				    <td class="col-sm-9"><a href="#WidgetCollectionModel" data-toggle="collapse" data-target="#WidgetCollectionModelInfo">WidgetCollection</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetsGetInfo -->
      </div><!-- .panel -->
	  <div class="panel panel-info">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Post</b></div>
			  <div class="col-sm-10">/api/widgets</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetsPostInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetsPostInfo">
			  <p>creates a new Widget in the collection</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				
				
				<h5>Body</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9"></th>
				  </tr>
				</thead>
				<tbody>
					<tr>
  				    <td class="col-sm-3"><a href="#WidgetRequestModel">WidgetRequest</a></td>
  				    <td class="col-sm-9"></td>
					</tr>
				</tbody>
			  </table>
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">201</td>
				    <td class="col-sm-9"><a href="#WidgetModel" data-toggle="collapse" data-target="#WidgetModelInfo">Widget</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetsPostInfo -->
      </div><!-- .panel -->
	  
	  <div class="panel panel-success">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Get</b></div>
			  <div class="col-sm-10">/api/widgets/{{id}}</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetGetInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetGetInfo">
			  <p>returns a the requested Widget</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				<h5>URI Parameters</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">id</td>
				    <td class="col-sm-9">string</td>
				  </tr></tbody></table>
				
				
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">200</td>
				    <td class="col-sm-9"><a href="#WidgetModel" data-toggle="collapse" data-target="#WidgetModelInfo">Widget</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetGetInfo -->
      </div><!-- .panel -->
	  <div class="panel panel-warning">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Put</b></div>
			  <div class="col-sm-10">/api/widgets/{{id}}</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetPutInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetPutInfo">
			  <p>replaces the identified Widget with the new Widget</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				<h5>URI Parameters</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">id</td>
				    <td class="col-sm-9">string</td>
				  </tr></tbody></table>
				
				<h5>Body</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9"></th>
				  </tr>
				</thead>
				<tbody>
					<tr>
  				    <td class="col-sm-3"><a href="#WidgetRequestModel">WidgetRequest</a></td>
  				    <td class="col-sm-9"></td>
					</tr>
				</tbody>
			  </table>
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">200</td>
				    <td class="col-sm-9"><a href="#WidgetModel" data-toggle="collapse" data-target="#WidgetModelInfo">Widget</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetPutInfo -->
      </div><!-- .panel -->
	  <div class="panel panel-warning">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Patch</b></div>
			  <div class="col-sm-10">/api/widgets/{{id}}</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetPatchInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetPatchInfo">
			  <p>updates fields on the identified Widget</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				<h5>URI Parameters</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">id</td>
				    <td class="col-sm-9">string</td>
				  </tr></tbody></table>
				
				<h5>Body</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9"></th>
				  </tr>
				</thead>
				<tbody>
					<tr>
  				    <td class="col-sm-3"><a href="#WidgetPatchModel">WidgetPatch</a></td>
  				    <td class="col-sm-9"></td>
					</tr>
				</tbody>
			  </table>
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">200</td>
				    <td class="col-sm-9"><a href="#WidgetModel" data-toggle="collapse" data-target="#WidgetModelInfo">Widget</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetPatchInfo -->
      </div><!-- .panel -->
	  <div class="panel panel-danger">
		  <div class="panel-heading">
			<div class="container">
			  <div class="col-sm-1"><b>Delete</b></div>
			  <div class="col-sm-10">/api/widgets/{{id}}</div>
			  <div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetDeleteInfo"><span class="glyphicon glyphicon-lock"></span></span></div>
			</div><!-- .container -->
		  </div>
		  <div class="panel-body collapse" id="WidgetDeleteInfo">
			  <p>deletes the identified Widget from the collection</p>
			  <h5>Headers</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">Authorization</td>
				    <td class="col-sm-9">Standard HTTP Basic Auth</td>
				  </tr>
				</tbody>
			  </table>
				<h5>URI Parameters</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Name</th>
				    <th class="col-sm-9">Description</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">id</td>
				    <td class="col-sm-9">string</td>
				  </tr></tbody></table>
				
				
			  <h5>Response</h5>
			  <table class="table">
				<thead>
				  <tr>
				    <th class="col-sm-3">Code</th>
				    <th class="col-sm-9">Model</th>
				  </tr>
				</thead>
				<tbody>
				  <tr>
				    <td class="col-sm-3">204</td>
				    <td class="col-sm-9"><a href="#WidgetModel" data-toggle="collapse" data-target="#WidgetModelInfo">Widget</a></td>
				  </tr>
				</tbody>
			  </table>
		  </div><!-- #WidgetDeleteInfo -->
      </div><!-- .panel -->
    </div><!-- .container -->

    <div class="container">
	  <h2>Models</h2>
	  
	  <div class="panel panel-default" id="WidgetModel">
		<div class="panel-heading">
          <div class="container">
            <div class="col-sm-11">Widget</div>
			<div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetModelInfo"><span class="glyphicon glyphicon-sort"></span></span></div>
          </div>
        </div>
		<div class="panel-body collapse" id="WidgetModelInfo">
		  <p>represent simple Widget for the example</p>
          <table class="table">
            <thead>
              <tr>
                <th class="col-sm-3">Name</th>
                <th class="col-sm-9">Type</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="col-sm-3">id</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">serial_number</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">description</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">created_on</td>
                <td class="col-sm-9">datetime</td>
			  </tr>
              <tr>
                <td class="col-sm-3">updated_on</td>
                <td class="col-sm-9">datetime</td>
			  </tr>
			</tbody>
          </table>
	    </div><!-- #WidgetModelInfo -->
	  </div><!-- #WidgetModel -->
	  <div class="panel panel-default" id="WidgetRequestModel">
		<div class="panel-heading">
          <div class="container">
            <div class="col-sm-11">WidgetRequest</div>
			<div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetRequestModelInfo"><span class="glyphicon glyphicon-sort"></span></span></div>
          </div>
        </div>
		<div class="panel-body collapse" id="WidgetRequestModelInfo">
		  <p>is the allowed input for a Widget (POST, PUT)</p>
          <table class="table">
            <thead>
              <tr>
                <th class="col-sm-3">Name</th>
                <th class="col-sm-9">Type</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="col-sm-3">serial_number</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">description</td>
                <td class="col-sm-9">string</td>
			  </tr>
			</tbody>
          </table>
	    </div><!-- #WidgetRequestModelInfo -->
	  </div><!-- #WidgetRequestModel -->
	  <div class="panel panel-default" id="WidgetPatchModel">
		<div class="panel-heading">
          <div class="container">
            <div class="col-sm-11">WidgetPatch</div>
			<div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetPatchModelInfo"><span class="glyphicon glyphicon-sort"></span></span></div>
          </div>
        </div>
		<div class="panel-body collapse" id="WidgetPatchModelInfo">
		  <p>is the allowed input for a Widget (PATCH)</p>
          <table class="table">
            <thead>
              <tr>
                <th class="col-sm-3">Name</th>
                <th class="col-sm-9">Type</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="col-sm-3">description</td>
                <td class="col-sm-9">string</td>
			  </tr>
			</tbody>
          </table>
	    </div><!-- #WidgetPatchModelInfo -->
	  </div><!-- #WidgetPatchModel -->
	  <div class="panel panel-default" id="CollectionModel">
		<div class="panel-heading">
          <div class="container">
            <div class="col-sm-11">Collection</div>
			<div class="col-sm-1"><span data-toggle="collapse" data-target="#CollectionModelInfo"><span class="glyphicon glyphicon-sort"></span></span></div>
          </div>
        </div>
		<div class="panel-body collapse" id="CollectionModelInfo">
		  <p>is the model for returning a Collection of objects in the API</p>
          <table class="table">
            <thead>
              <tr>
                <th class="col-sm-3">Name</th>
                <th class="col-sm-9">Type</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="col-sm-3">next</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">previous</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">uri</td>
                <td class="col-sm-9">string</td>
			  </tr>
			</tbody>
          </table>
	    </div><!-- #CollectionModelInfo -->
	  </div><!-- #CollectionModel -->
      
	  <div class="panel panel-default" id="WidgetCollectionModel">
		<div class="panel-heading">
          <div class="container">
            <div class="col-sm-11">WidgetCollection</div>
			<div class="col-sm-1"><span data-toggle="collapse" data-target="#WidgetCollectionModelInfo"><span class="glyphicon glyphicon-sort"></span></span></div>
          </div>
        </div>
		<div class="panel-body collapse" id="WidgetCollectionModelInfo">
		  <p>WidgetCollection is a paginated collection of Widget models</p>
          <table class="table">
            <thead>
              <tr>
                <th class="col-sm-3">Name</th>
                <th class="col-sm-9">Type</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="col-sm-3">next</td>
                <td class="col-sm-9">string</td>
              </tr>
              <tr>
                <td class="col-sm-3">previous</td>
                <td class="col-sm-9">string</td>
              </tr>
              <tr>
                <td class="col-sm-3">uri</td>
                <td class="col-sm-9">string</td>
			  </tr>
              <tr>
                <td class="col-sm-3">items</td>
                <td class="col-sm-9"><a href="#WidgetModel">[]Widget</a></td>
              </tr>
			</tbody>
          </table>
	    </div><!-- #WidgetCollectionModelInfo -->
	  </div><!-- #WidgetCollectionModel -->
  </div><!-- .contrainer -->

	<!-- Optional JavaScript -->
	<!-- jQuery first, then Bootstrap JS -->
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
  </body>
</html>`
